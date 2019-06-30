package elf_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kelfa/elf"
)

func TestWriteAll(t *testing.T) {
	b := &bytes.Buffer{}
	w := elf.NewWriter(b, []string{"A", "B", "C"})
	err := w.WriteAll(
		[]map[string]string{
			map[string]string{
				"C": "c",
				"A": "a",
				"B": "b",
			},
		})
	if err != nil {
		t.Errorf("Unexpected error:\ngot  %v\nwant %v", err, nil)
	}
	out := b.String()
	fmt.Println(out)
}
