package elf_test

import (
	"bytes"
	"testing"

	"github.com/kelfa/kelfa/elf"
)

func TestWriteAll(t *testing.T) {
	tests := []struct {
		Fields         []string
		Entries        []map[string]string
		UseCRLF        bool
		ExpectedError  error
		ExpectedOutput string
	}{
		// Parameters inverted
		{
			Fields: []string{"A", "B", "C"},
			Entries: []map[string]string{
				{
					"C": "c",
					"A": "a",
					"B": "b",
				},
			},
			ExpectedError:  nil,
			ExpectedOutput: "#Version: 1.0\n#Fields: A\tB\tC\na\tb\tc\n",
		},
		// Missing parameter
		{
			Fields: []string{"A", "B", "C"},
			Entries: []map[string]string{
				{
					"A": "a",
					"B": "b",
				},
			},
			ExpectedError:  nil,
			ExpectedOutput: "#Version: 1.0\n#Fields: A\tB\tC\na\tb\t-\n",
		},
		// CRLF
		{
			Fields:  []string{"A", "B", "C"},
			UseCRLF: true,
			Entries: []map[string]string{
				{
					"A": "a",
					"B": "b",
					"C": "c",
				},
			},
			ExpectedError:  nil,
			ExpectedOutput: "#Version: 1.0\r\n#Fields: A\tB\tC\r\na\tb\tc\r\n",
		},
	}

	for _, tc := range tests {
		b := &bytes.Buffer{}
		w := elf.NewWriter(b, tc.Fields)
		w.UseCRLF = tc.UseCRLF
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
