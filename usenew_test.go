package usenew

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		dir string
	}{
		{dir: "a"},
		{dir: "b"},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			runTests(t, Analyzer, test.dir)
		})
	}
}

func runTests(t *testing.T, a *analysis.Analyzer, dir string, patterns ...string) []*analysistest.Result {
	t.Helper()

	tempDir := t.TempDir()

	// Needs to be run outside testdata.
	err := os.CopyFS(tempDir, os.DirFS(filepath.Join(analysistest.TestData(), "src")))
	if err != nil {
		t.Fatal(err)
	}

	// NOTE: analysistest does not yet support modules;
	// see https://github.com/golang/go/issues/37054 for details.

	srcPath := filepath.Join(tempDir, filepath.FromSlash(dir))

	t.Chdir(srcPath)

	output, err := exec.CommandContext(t.Context(), "go", "mod", "vendor").CombinedOutput()
	if err != nil {
		t.Log(string(output))
		t.Fatal(err)
	}

	return analysistest.Run(t, srcPath, a, patterns...)
}
