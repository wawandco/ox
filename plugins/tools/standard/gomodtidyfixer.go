package standard

import (
	"context"

	"github.com/wawandco/ox/internal/source"
)

// GoModTidyFixer is a fixer that runs `go mod tidy`.
type GoModTidyFixer struct{}

func (ef GoModTidyFixer) Name() string {
	return "standard/modtidy"
}

func (ef GoModTidyFixer) Fix(ctx context.Context, root string, args []string) error {
	return source.RunModTidy(root)
}
