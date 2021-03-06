package standard

import (
	"context"
	_ "embed"
	"path/filepath"

	"github.com/wawandco/ox/internal/runtime"
	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/plugins/base/new"
)

var (
	//go:embed templates/go.mod.tmpl
	goModTemplate string
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "standard/initializer"
}

// Initialize the go module
func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	opts := struct {
		OxVersion  string
		ModuleName string
	}{
		OxVersion:  runtime.Version,
		ModuleName: options.Module,
	}

	err := source.Build(filepath.Join(options.Folder, "go.mod"), goModTemplate, opts)
	return err
}
