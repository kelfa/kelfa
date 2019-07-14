package filepath_test

import (
	"testing"

	"go.kelfa.io/kelfa/pkg/filepath"
)

type ExtTC struct {
	Path           string
	DefExt         string
	ExpectedOutput string
}

var ExtTestCases = []ExtTC{
	ExtTC{Path: "/home/test/file.txt", DefExt: "html", ExpectedOutput: "txt"}, // Standard file
	ExtTC{Path: "/home/test/file", DefExt: "html", ExpectedOutput: "html"},    // File without extention
	ExtTC{Path: "/home/test/", DefExt: "html", ExpectedOutput: "html"},        // Folder with tailing slash
	ExtTC{Path: "/home/test", DefExt: "html", ExpectedOutput: "html"},         // Folder with no tailing slash
	ExtTC{Path: "/ho.me/test", DefExt: "html", ExpectedOutput: "html"},        // Folder with no tailing slash
}

func TestExt(t *testing.T) {
	for _, tc := range ExtTestCases {
		if filepath.Ext(tc.Path, tc.DefExt) != tc.ExpectedOutput {
			t.Fatalf("for %s, expecting %s, got %s", tc.Path, tc.ExpectedOutput, filepath.Ext(tc.Path, tc.DefExt))
		}
	}
}
