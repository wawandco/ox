package action

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/internal/source"
)

var (

	//go:embed templates
	templates embed.FS

	files = map[string]string{
		"actions_test.go.tmpl": filepath.Join("app", "actions", "actions_test.go"),
		"actions.go.tmpl":      filepath.Join("app", "actions", "actions.go"),
		"home.go.tmpl":         filepath.Join("app", "actions", "home", "home.go"),
	}

	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context) error {
	m := ctx.Value("module")
	if m == nil {
		return ErrIncompleteArgs
	}

	f := ctx.Value("folder")
	if f == nil {
		return ErrIncompleteArgs
	}

	entries, err := templates.ReadDir("templates")
	if err != nil {
		return err
	}

	for _, e := range entries {

		if e.IsDir() {
			continue
		}

		bt, err := fs.ReadFile(templates, filepath.Join("templates", e.Name()))
		if err != nil {
			return err
		}

		template := string(bt)
		result := files[e.Name()]
		if result == "" {
			continue
		}

		err = source.Build(filepath.Join(f.(string), result), template, m)
		if err != nil {
			return err
		}
	}

	fmt.Printf("[info] Created app/actions/actions.go\n")
	fmt.Printf("[info] Created app/actions/actions_test.go\n")
	fmt.Printf("[info] Created app/actions/home/home.go\n")

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
