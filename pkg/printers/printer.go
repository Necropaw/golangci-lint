package printers

import (
	"context"

	"github.com/necropaw/golangci-lint/pkg/result"
)

type Printer interface {
	Print(ctx context.Context, issues []result.Issue) error
}
