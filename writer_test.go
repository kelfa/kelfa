package elf_test

import (
	"bytes"
	"testing"

	"github.com/kelfa/elf"
)

type WriterTestCase struct {
	Fields         []string
	Entries        []map[string]string
	ExpectedError  error
	ExpectedOutput string
}

var writerTestCases = []WriterTestCase{
	WriterTestCase{
		Fields: []string{"A", "B", "C"},
		Entries: []map[string]string{
			map[string]string{
				"C": "c",
				"A": "a",
				"B": "b",
			},
		},
		ExpectedError:  nil,
		ExpectedOutput: "#Version: 1.0\n#Fields: A\tB\tC\na\tb\tc\n",
	},
}

func TestWriteAll(t *testing.T) {
	for _, tc := range writerTestCases {
		b := &bytes.Buffer{}
		w := elf.NewWriter(b, tc.Fields)
		err := w.WriteAll(tc.Entries)
		if err != tc.ExpectedError {
			t.Errorf("Unexpected error: got %v wanted %v", err, tc.ExpectedError)
		}
		if b.String() != tc.ExpectedOutput {
			t.Errorf("Unexpected output: got:\n%v\nwanted:\n%v", b.String(), tc.ExpectedOutput)
			t.Errorf("Unexpected output: got:\n%v\nwanted:\n%v", b.Bytes(), []byte(tc.ExpectedOutput))
		}
	}
}
