package processors

import (
	"testing"

	"github.com/necropaw/golangci-lint/pkg/config"
	"github.com/necropaw/golangci-lint/pkg/logutils"
)

func TestMaxFromLinter(t *testing.T) {
	p := NewMaxFromLinter(1, logutils.NewStderrLog(logutils.DebugKeyEmpty), &config.Config{})
	gosimple := newFromLinterIssue("gosimple")
	gofmt := newFromLinterIssue("gofmt")
	processAssertSame(t, p, gosimple)  // ok
	processAssertSame(t, p, gofmt)     // ok: another
	processAssertEmpty(t, p, gosimple) // skip
}
