package git

import (
	"context"
	"embed"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

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
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	keeps := []string{
		"migrations",
		"public",
	}

	for _, k := range keeps {
		err := os.MkdirAll(filepath.Join(options.Folder, k), 0777)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(options.Folder, k, ".gitkeep"), []byte{}, 0777)
		if err == nil {
			continue
		}

		return err
	}

	files := []struct {
		path     string
		template string
	}{
		{filepath.Join(options.Folder, ".gitignore"), "templates/dot-gitignore.tmpl"},
	}

	for _, f := range files {
		content, err := templates.ReadFile(f.template)
		if err != nil {
			return err
		}

		err = source.Build(f.path, string(content), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
