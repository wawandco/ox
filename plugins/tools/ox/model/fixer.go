package model

import (
	"context"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/source"
)

// Fixer
type Fixer struct{}

func (ef Fixer) Name() string {
	return "models/fixer"
}

func (ef Fixer) Fix(ctx context.Context, root string, args []string) error {
	tmpl, err := Templates.ReadFile("templates/models.go.tmpl")
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
