// Copyright 2019 The Kelfe authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package elf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// Writer writes records to an ELF encoded file.
type Writer struct {
	Fields    []string
	Version   string
	Software  string
	StartDate string
	EndDate   string
	Date      string
	Remark    string
	Separator rune // Field delimiter (set to '\t' by NewWriter)
	UseCRLF   bool // True to use \r\n as the line terminator
	w         *bufio.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer, fields []string) *Writer {
	return &Writer{
		Separator: '\t',
		w:         bufio.NewWriter(w),
		Version:   "1.0",
		Fields:    fields,
	}
}

// WriteAll writes multiple elf records to w using Write and then calls Flush.
func (w *Writer) WriteAll(records []map[string]string) error {
	if err := w.WriteHeaders(); err != nil {
		return err
	}
	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			return err
		}
	}
	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush() // #nosec
}

// Error reports any error that has occurred during a previous Write or Flush.
func (w *Writer) Error() error {
	_, err := w.w.Write(nil)
	return err
}

func validDelim(r rune) bool {
	return r == ' ' || r == '\t'
}

// WriteHeaders write the ELF file header lines
func (w *Writer) WriteHeaders() error {
	if len(w.Fields) == 0 {
		return errors.New("the Fields field needs to be populated")
	}
	if !validDelim(w.Separator) {
		return errors.New("invalid field or comment delimiter")
	}

	if err := w.writeVersion(); err != nil {
		return err
	}
	if err := w.writeFields(); err != nil {
		return err
	}
	return nil
}

func (w *Writer) writeVersion() error {
	_, err := w.w.WriteString(fmt.Sprintf("#Version: %s", w.Version))
	if err != nil {
		return err
	}
	if w.UseCRLF {
		_, err = w.w.WriteString("\r\n")
	} else {
		err = w.w.WriteByte('\n')
	}
	return err
}

func (w *Writer) writeFields() error {
	_, err := w.w.WriteString("#Fields: ")
	if err != nil {
		return err
	}
	for c, field := range w.Fields {
		if c > 0 {
			_, err = w.w.WriteRune(w.Separator)
			if err != nil {
				return err
			}
		}
		_, err = w.w.WriteString(field)
		if err != nil {
			return err
		}
	}

	if w.UseCRLF {
		_, err = w.w.WriteString("\r\n")
	} else {
		err = w.w.WriteByte('\n')
	}
	return err
}

func (w *Writer) Write(record map[string]string) error {
	var err error
	if len(w.Fields) == 0 {
		return errors.New("the Fields field needs to be populated")
	}
	if !validDelim(w.Separator) {
		return errors.New("invalid field or comment delimiter")
	}

	// Pointless to write an empty line
	if len(record) == 0 {
		return nil
	}

	for c, field := range w.Fields {
		if c > 0 {
			_, err = w.w.WriteRune(w.Separator)
			if err != nil {
				return err
			}
		}
		if val, ok := record[field]; ok {
			_, err = w.w.WriteString(val)
			if err != nil {
				return err
			}
		} else {
			_, err = w.w.WriteString("-")
			if err != nil {
				return err
			}
		}
	}

	if w.UseCRLF {
		_, err = w.w.WriteString("\r\n")
	} else {
		err = w.w.WriteByte('\n')
	}
	return err
}
