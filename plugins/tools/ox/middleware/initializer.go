package middleware

import (
	"context"
	_ "embed"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/plugins/base/new"
)

var (
	//go:embed middleware.go.tmpl
	middlewareTemplate string
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "middleware/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {

	filename := filepath.Join(options.Folder, "app", "middleware", "middleware.go")
	err := source.Build(filename, middlewareTemplate, options.Module)
	return err
}
