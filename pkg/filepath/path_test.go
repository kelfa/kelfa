package filepath_test

import (
	"testing"

	"go.kelfa.io/kelfa/pkg/filepath"
)

func TestPath(t *testing.T) {
	tests := []struct {
		Path           string
		ExpectedOutput string
	}{
		{Path: "/home/test/file.txt", ExpectedOutput: "/home/test"}, // Standard file
		{Path: "/home/test/file", ExpectedOutput: "/home/test"},     // File without extension
		{Path: "/home/test/", ExpectedOutput: "/home/test"},         // Folder with tailing slash
		{Path: "/home/test", ExpectedOutput: "/home"},               // Folder with no tailing slash
	}

	for _, tc := range tests {
		if filepath.Path(tc.Path) != tc.ExpectedOutput {
			t.Fatalf("expecting %s, got %s", tc.ExpectedOutput, filepath.Path(tc.Path))
		}
	}
}
