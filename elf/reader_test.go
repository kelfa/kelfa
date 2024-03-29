package elf_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/kelfa/kelfa/elf"
)

func TestReadAll(t *testing.T) {
	tests := []struct {
		ELF             string
		UseCRLF         bool
		ExpectedError   error
		ExpectedVersion string
		ExpectedOutput  []map[string]string
		ExpectedFields  []string
	}{
		{
			ELF: "#Version: 1.0\n" +
				"#Date: 12-Jan-1996 00:00:00\n" +
				"#Fields: time cs-method cs-uri\n" +
				"00:34:23 GET /foo/bar.html\n" +
				"12:21:16 GET /foo/bar.html\n" +
				"12:45:52 GET /foo/bar.html\n" +
				"12:57:34 GET /foo/bar.html",
			ExpectedError:   nil,
			ExpectedVersion: "1.0",
			ExpectedFields:  []string{"time", "cs-method", "cs-uri"},
			ExpectedOutput: []map[string]string{
				{
					"time":      "00:34:23",
					"cs-method": "GET",
					"cs-uri":    "/foo/bar.html",
				},
				{
					"time":      "12:21:16",
					"cs-method": "GET",
					"cs-uri":    "/foo/bar.html",
				},
				{
					"time":      "12:45:52",
					"cs-method": "GET",
					"cs-uri":    "/foo/bar.html",
				},
				{
					"time":      "12:57:34",
					"cs-method": "GET",
					"cs-uri":    "/foo/bar.html",
				},
			},
		},
	}

	for _, tc := range tests {
		r := elf.NewReader(strings.NewReader(tc.ELF))
		d, err := r.ReadAll()
		if err != tc.ExpectedError {
			t.Errorf("Unexpected error: got %v wanted %v", err, tc.ExpectedError)
		}
		if !reflect.DeepEqual(d, tc.ExpectedOutput) {
			t.Errorf("Unexpected output: got:\n%v\nwanted:\n%v", d, tc.ExpectedOutput)
		}
		if r.Version != tc.ExpectedVersion {
			t.Errorf("Unexpected version: got %v wanted %v", r.Version, tc.ExpectedVersion)
		}
		if !reflect.DeepEqual(r.Fields, tc.ExpectedFields) {
			t.Errorf("Unexpected fields: got %v wanted %v", r.Fields, tc.ExpectedFields)
		}
	}
}
