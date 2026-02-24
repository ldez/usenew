package usenew

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		dir string
	}{
		{dir: "a"},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			analysistest.Run(t, filepath.Join(analysistest.TestData(), "src", test.dir), Analyzer)
		})
	}
}
