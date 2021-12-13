package embedded

import (
	"context"
	"os"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/plugins/base/new"
)

// Fixer
type Fixer struct{}

func (ef Fixer) Name() string {
	return "embedded/fixer"
}

func (ef Fixer) Fix(ctx context.Context, root string, args []string) error {
	err := os.Remove(filepath.Join(root, "embed.go"))
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	name, err := info.BuildName()
	if err != nil {
		return err
	}

	ini := &Initializer{}
	err = ini.Initialize(ctx, new.Options{
		Folder: root,
		Name:   name,
	})

	return err
}
