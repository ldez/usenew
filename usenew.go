package usenew

import (
	"go/ast"
	"go/build"
	"go/types"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	go126   = 126
	goDevel = 666
)

var Analyzer = &analysis.Analyzer{
	Name:     "usenew",
	Doc:      "Find calls to functions that can be replaced with the built-in 'new' function.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	goVersion := getGoVersion(pass)
	if goVersion < go126 {
		return nil, nil
	}

	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		if v, ok := call.Fun.(*ast.Ident); ok {
			if v.Name == "new" {
				return
			}
		}

		typ := pass.TypesInfo.TypeOf(call.Fun)
		if typ == nil {
			return
		}

		sig, ok := typ.(*types.Signature)
		if !ok {
			return
		}

		if sig.Params().Len() != 1 || sig.Results().Len() != 1 {
			return
		}

		paramType := sig.Params().At(0).Type()
		resultType := sig.Results().At(0).Type()

		ptrType, ok := resultType.(*types.Pointer)
		if !ok {
			return
		}

		if types.Identical(ptrType.Elem(), paramType) {
			pass.Report(analysis.Diagnostic{
				Pos:            call.Pos(),
				End:            call.End(),
				Message:        "This call can be replaced with the built-in 'new' function.",
				SuggestedFixes: nil,
			})
		}
	})

	return nil, nil
}

func getGoVersion(pass *analysis.Pass) int {
	// Prior to go1.22, versions.FileVersion returns only the toolchain version,
	// which is of no use to us,
	// so disable this analyzer on earlier versions.
	if !slices.Contains(build.Default.ReleaseTags, "go1.22") {
		return 0 // false
	}

	pkgVersion := pass.Pkg.GoVersion()
	if pkgVersion == "" {
		// Empty means Go devel.
		return goDevel // true
	}

	raw := strings.TrimPrefix(pkgVersion, "go")

	// prerelease version (go1.24rc1)
	idx := strings.IndexFunc(raw, func(r rune) bool {
		return (r < '0' || r > '9') && r != '.'
	})

	if idx != -1 {
		raw = raw[:idx]
	}

	vParts := strings.Split(raw, ".")

	v, err := strconv.Atoi(strings.Join(vParts[:2], ""))
	if err != nil {
		v = 116
	}

	return v
}
