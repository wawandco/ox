package ox

import (
	"context"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/plugins/tools/buffalo/model"
)

// ModelsFixer
type ModelsFixer struct{}

func (ef ModelsFixer) Name() string {
	return "ox/fixer/models"
}

func (ef ModelsFixer) Fix(ctx context.Context, root string, args []string) error {
	tmpl, err := model.Templates.ReadFile("templates/models.go.tmpl")
	if err != nil {
		return err
	}

	mod := info.ModuleName()
	if err != nil {
		return err
	}

	filename := filepath.Join(root, "app", "models", "models.go")
	err = source.Build(filename, string(tmpl), mod)
	if err != nil {
		return err
	}

	return err
}
