package tasks

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
	"github.com/wawandco/oxpecker/internal/source"
)

type Generator struct {
	name     string
	filename string
	dir      string
}

func (g Generator) Name() string {
	return "task"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate task [name]`")
	}

	dirPath := filepath.Join(root, "app", "tasks")
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
		return errors.Errorf("Task file already exists")
	}

	if err := g.createTaskFile(args); err != nil {
		return errors.Wrap(err, "creating action file")
	}

	fmt.Printf("[info] task generated in: \n-- app/tasks/%s.go\n", g.name)

	return nil
}

func (g Generator) createTaskFile(args []string) error {
	path := filepath.Join(g.dir, g.filename+".go")
	err := source.Build(path, taskTemplate, g.name)

	return errors.Wrap(err, "parsing new template error")
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
