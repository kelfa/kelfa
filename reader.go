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

type Reader struct {
	Fields    []string
	Version   string
	Software  string
	StartDate string
	EndDate   string
	Date      string
	Remark    string

	r         *bufio.Reader
	rawBuffer []byte
	numLine   int
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

func (r *Reader) Read() (map[string]string, error) {
	record, err := r.readRecord()
	return record, err
}

func (r *Reader) ReadAll() ([]map[string]string, error) {
	var records []map[string]string
	for {
		record, err := r.readRecord()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
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
		return nil, fmt.Errorf("was expecting %v columns while line %v has %v", len(r.Fields), r.numLine, len(fs))
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
