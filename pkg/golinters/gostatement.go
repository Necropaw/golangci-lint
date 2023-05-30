package golinters

import (
	"github.com/necropaw/golangci-lint/pkg/golinters/goanalysis"

	"github.com/fpuc/gostatement"
	"golang.org/x/tools/go/analysis"
)

func NewGostatement() *goanalysis.Linter {

	return goanalysis.NewLinter(
		"gostatement",
		"gostatement is an analyzer checking for occurrence of `go` statements",
		[]*analysis.Analyzer{gostatement.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)

}
