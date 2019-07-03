// Copyright 2019 The Kelfe authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package elf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	ErrFieldCount = errors.New("wrong number of fields")
)

// A Reader reads records from an Extended Log Format file.
//
// As returned by NewReader, a Reader expects input conforming to the W3C ELF.
// The exported fields are populated by the Read and the ReadAll functions.
//
// The Reader manages both \r\n and \n sequences.
type Reader struct {
	// List of fields as for the Field descriptor.
	// This field gets populated by reader during the read.
	Fields []string
	// Version of the ELF format used, as for the Version descriptor.
	// This field gets populated by reader during the read.
	Version string
	// Software that created the ELF file, as for the Software descriptor.
	// This field gets populated by reader during the read.
	Software string
	// StartDate of the oldest entry, as for the StartDate descriptor.
	// This field gets populated by reader during the read.
	StartDate string
	// EndDate of the oldest entry, as for the EndDate descriptor.
	// This field gets populated by reader during the read.
	EndDate string
	// Date of the oldest entry, as for the Date descriptor.
	// This field gets populated by reader during the read.
	Date string
	// Remarks, as for the Remarks descriptor.
	// This field gets populated by reader during the read.
	Remark string

	r         *bufio.Reader
	rawBuffer []byte
	numLine   int
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

// Read reads one record (a slice of fields) from r.
// If the record has an unexpected number of fields,
// Read returns the error ErrFieldCount.
//
// Read always returns either a non-nil record or a non-nil error, but not both.
// The only exception is if the line is a comment or empty, then both the
// record and the error will be nil.
//
// If there is no data left to be read, Read returns nil, io.EOF.
func (r *Reader) Read() (record map[string]string, err error) {
	record, err = r.readRecord()
	return record, err
}

// ReadAll reads all the remaining records from r.
// Each record is a map of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func (r *Reader) ReadAll() (records []map[string]string, err error) {
	for {
		record, err := r.readRecord()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		if record != nil {
			records = append(records, record)
		}
	}
}

func (r *Reader) readRecord() (map[string]string, error) {
	bs, err := r.readLine()
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, nil
	}
	// If it starts with a # it is either a Directive or a comment
	if bs[0] == '#' {
		r.identifyDirectives(bs)
		return nil, nil
	}
	if r.Version != "1.0" {
		return nil, fmt.Errorf("elf package only supports elf version 1.0")
	}
	if len(r.Fields) == 0 {
		return nil, errors.New("no Fields directive found")
	}
	fs := strings.Fields(string(bs))
	if len(fs) != len(r.Fields) {
		return nil, ErrFieldCount
	}
	record := make(map[string]string)
	for i, f := range fs {
		record[r.Fields[i]] = f
	}
	return record, nil
}

func (r *Reader) identifyDirectives(bs []byte) {
	if dir := getDirective(bs, []byte("#Version: ")); len(dir) > 0 {
		r.Version = string(dir)
	}
	if dir := getDirective(bs, []byte("#Fields: ")); len(dir) > 0 {
		r.Fields = strings.Fields(string(dir))
	}
	if dir := getDirective(bs, []byte("#Software: ")); len(dir) > 0 {
		r.Software = string(dir)
	}
	if dir := getDirective(bs, []byte("#Start-Date: ")); len(dir) > 0 {
		r.StartDate = string(dir)
	}
	if dir := getDirective(bs, []byte("#End-Date: ")); len(dir) > 0 {
		r.EndDate = string(dir)
	}
	if dir := getDirective(bs, []byte("#Date: ")); len(dir) > 0 {
		r.Date = string(dir)
	}
	if dir := getDirective(bs, []byte("#Remark: ")); len(dir) > 0 {
		r.Remark = string(dir)
	}
	return
}

func getDirective(bs []byte, directive []byte) []byte {
	if len(bs) < len(directive) {
		return nil
	}
	for i, v := range directive {
		if v != bs[i] {
			return nil
		}
	}
	return bs[len(directive) : len(bs)-1]
}

func (r *Reader) readLine() ([]byte, error) {
	line, err := r.r.ReadSlice('\n')
	if err == bufio.ErrBufferFull {
		r.rawBuffer = append(r.rawBuffer[:0], line...)
		for err == bufio.ErrBufferFull {
			line, err = r.r.ReadSlice('\n')
			r.rawBuffer = append(r.rawBuffer, line...)
		}
		line = r.rawBuffer
	}
	if len(line) > 0 && err == io.EOF {
		err = nil
		// For backwards compatibility, drop trailing \r before EOF.
		if line[len(line)-1] == '\r' {
			line = line[:len(line)-1]
		}
	}
	r.numLine++
	// Normalize \r\n to \n on all input lines.
	if n := len(line); n >= 2 && line[n-2] == '\r' && line[n-1] == '\n' {
		line[n-2] = '\n'
		line = line[:n-1]
	}
	return line, err
}
