package model

import (
	"context"
	"os"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/plugins/base/new"
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	tmpl, err := Templates.ReadFile("templates/models.go.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(options.Folder, "app", "models", "models.go")
	err = source.Build(filename, string(tmpl), options.Module)
	if err != nil {
		return err
	}

	tmpl, err = Templates.ReadFile("templates/models_test.go.tmpl")
	if err != nil {
		return err
	}

	filename = filepath.Join(options.Folder, "app", "models", "models_test.go")
	err = os.WriteFile(filename, tmpl, 0777)
	if err != nil {
		return err
	}

	return nil
}
