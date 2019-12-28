package filepath_test

import (
	"testing"

	"go.kelfa.io/pkg/filepath"
)

func TestExt(t *testing.T) {
	tests := []struct {
		Path           string
		DefExt         string
		ExpectedOutput string
	}{
		{Path: "/home/test/file.txt", DefExt: "html", ExpectedOutput: "txt"}, // Standard file
		{Path: "/home/test/file", DefExt: "html", ExpectedOutput: "html"},    // File without extension
		{Path: "/home/test/", DefExt: "html", ExpectedOutput: "html"},        // Folder with tailing slash
		{Path: "/home/test", DefExt: "html", ExpectedOutput: "html"},         // Folder with no tailing slash
		{Path: "/ho.me/test", DefExt: "html", ExpectedOutput: "html"},        // Folder with no tailing slash
	}

	for _, tc := range tests {
		if filepath.Ext(tc.Path, tc.DefExt) != tc.ExpectedOutput {
			t.Fatalf("for %s, expecting %s, got %s", tc.Path, tc.ExpectedOutput, filepath.Ext(tc.Path, tc.DefExt))
		}
	}
}
