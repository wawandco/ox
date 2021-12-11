package ox

import (
	"context"
	"os"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/plugins/base/new"
	"github.com/wawandco/ox/plugins/tools/buffalo/embedded"
)

// EmbedFixer
type EmbedFixer struct{}

func (ef EmbedFixer) Name() string {
	return "ox/fixer/embed"
}

func (ef EmbedFixer) Fix(ctx context.Context, root string, args []string) error {
	err := os.Remove(filepath.Join(root, "embed.go"))
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	name, err := info.BuildName()
	if err != nil {
		return err
	}

	ini := &embedded.Initializer{}
	err = ini.Initialize(ctx, new.Options{
		Folder: root,
		Name:   name,
	})

	return err
}
