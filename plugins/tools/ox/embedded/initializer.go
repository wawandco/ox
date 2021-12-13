package embedded

import (
	"context"
	"embed"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/plugins/base/new"
)

var (

	//go:embed templates
	templates embed.FS
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "embedded/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	files := map[string]string{
		"templates/rootname.go.tmpl":   filepath.Join(options.Folder, options.Name+".go"),
		"templates/templates.go.tmpl":  filepath.Join(options.Folder, "app", "templates", "templates.go"),
		"templates/public.go.tmpl":     filepath.Join(options.Folder, "public", "public.go"),
		"templates/config.go.tmpl":     filepath.Join(options.Folder, "config", "config.go"),
		"templates/migrations.go.tmpl": filepath.Join(options.Folder, "migrations", "migrations.go"),
	}

	for k, path := range files {
		content, err := templates.ReadFile(k)
		if err != nil {
			return err
		}

		err = source.Build(path, string(content), options.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
