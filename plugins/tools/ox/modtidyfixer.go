package ox

import (
	"context"

	"github.com/wawandco/ox/internal/source"
)

type ModTidyFixer struct{}

func (ef ModTidyFixer) Name() string {
	return "ox/fixer/modtidy"
}

func (ef ModTidyFixer) Fix(ctx context.Context, root string, args []string) error {
	return source.RunModTidy(root)
}
