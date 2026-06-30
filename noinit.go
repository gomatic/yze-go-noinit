// Package noinit provides a go/analysis analyzer that forbids a package-level
// func init, per the gomatic Go standard that favors explicit construction and
// dependency injection over implicit package initialization.
package noinit

import (
	"go/ast"

	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const message = "func init is not permitted; use explicit construction or dependency injection"

// Analyzer reports every package-level func init.
var Analyzer = &analysis.Analyzer{
	Name:     "noinit",
	Doc:      "reports a package-level func init, which the gomatic Go standard forbids in favor of explicit construction / dependency injection",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "noinit",
	Categories: []goyze.Category{"patterns"},
	URL:        "https://docs.gomatic.dev/yze/noinit",
	Analyzer:   Analyzer,
}

// run reports every package-level func init in the analyzed package. The
// special init function has no receiver; a method named init (one with a
// receiver) is an ordinary method and is left unreported.
func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	insp.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
		if fn := n.(*ast.FuncDecl); fn.Recv == nil && fn.Name.Name == "init" {
			pass.Reportf(fn.Pos(), message)
		}
	})
	return nil, nil
}
