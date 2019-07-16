package filepath_test

import (
	"testing"

	"go.kelfa.io/kelfa/pkg/filepath"
)

type PathTC struct {
	Path           string
	ExpectedOutput string
}

var PathTestCases = []PathTC{
	{Path: "/home/test/file.txt", ExpectedOutput: "/home/test"}, // Standard file
	{Path: "/home/test/file", ExpectedOutput: "/home/test"},     // File without extension
	{Path: "/home/test/", ExpectedOutput: "/home/test"},         // Folder with tailing slash
	{Path: "/home/test", ExpectedOutput: "/home"},               // Folder with no tailing slash
}

func TestPath(t *testing.T) {
	for _, tc := range PathTestCases {
		if filepath.Path(tc.Path) != tc.ExpectedOutput {
			t.Fatalf("expecting %s, got %s", tc.ExpectedOutput, filepath.Path(tc.Path))
		}
	}
}
