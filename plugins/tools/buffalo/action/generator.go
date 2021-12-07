package action

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"errors"

	"github.com/gobuffalo/flect"
	"github.com/wawandco/ox/internal/source"
)

type Generator struct {
	name     string
	filename string
	dir      string
}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "buffalo/generate-action"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
	return "action"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("no name specified, please use `ox generate action [name]`")
	}

	dirPath := filepath.Join(root, "app", "actions")
	if !g.exists(dirPath) {
		err := os.MkdirAll(filepath.Dir(dirPath), 0755)
		if err != nil {
			return (err)
		}
	}

	g.name = flect.Singularize(args[2])
	g.filename = flect.Singularize(flect.Underscore(args[2]))
	g.dir = dirPath

	if g.exists(filepath.Join(g.dir, g.filename+".go")) {
		return errors.New("action file already exists")
	}

	if err := g.generateActionFiles(args[3:]); err != nil {
		return err
	}

	return nil
}

func (g Generator) generateActionFiles(args []string) error {
	if err := g.createActionFile(args); err != nil {
		return fmt.Errorf("creating action file: %w", err)
	}

	if err := g.createActionTestFile(); err != nil {
		return fmt.Errorf("creating action test file: %w", err)
	}

	return nil
}

func (g Generator) createActionFile(args []string) error {
	path := filepath.Join(g.dir, g.filename+".go")
	data := struct {
		Name string
	}{
		Name: g.name,
	}
	actionTemplate, err := g.callTemplate("action.go.tmpl")
	if err != nil {
		return fmt.Errorf("error calling template: %w", err)
	}

	err = source.Build(path, actionTemplate, data)
	if err != nil {
		return fmt.Errorf("error generating action: %w", err)
	}

	return nil
}

func (g Generator) createActionTestFile() error {
	path := filepath.Join(g.dir, g.filename+"_test.go")
	data := struct {
		Name string
	}{
		Name: g.name,
	}

	actionTestTemplate, err := g.callTemplate("action_test.go.tmpl")
	if err != nil {
		return fmt.Errorf("error calling template: %w", err)
	}

	err = source.Build(path, actionTestTemplate, data)
	if err != nil {
		return fmt.Errorf("error generating action: %w", err)
	}

	return nil
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func (g Generator) callTemplate(name string) (string, error) {
	bt, err := fs.ReadFile(templates, filepath.Join("templates", name))
	if err != nil {
		return "", nil
	}

	return string(bt), nil
}
