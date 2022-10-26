package cli

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/plugins/base/build"
	"github.com/wawandco/ox/plugins/base/dev"
	"github.com/wawandco/ox/plugins/base/generate"
	"github.com/wawandco/ox/plugins/base/help"
	"github.com/wawandco/ox/plugins/base/test"
	"github.com/wawandco/ox/plugins/base/version"
	"github.com/wawandco/ox/plugins/core"
	"github.com/wawandco/ox/plugins/tools/db"
)

func Test_CliTestingAliaser(t *testing.T) {
	plugins := []core.Plugin{
		&generate.Command{},
		&build.Command{},
		&dev.Command{},
		&version.Command{},
		&help.Command{},
		&test.Command{},
		&db.Command{},
	}

	c := &cli{
		plugins,
	}

	tcases := []struct {
		commandAlias string
		nameExpected string
	}{
		{"g", "generate"},
		{"b", "build"},
		{"d", "dev"},
		{"v", "version"},
		{"h", "help"},
		{"db", "database"},
		{"t", "test"},
	}

	for _, ca := range tcases {
		command := c.findCommand(ca.commandAlias)
		if command == nil {
			t.Errorf("Command %s not found", ca.commandAlias)
			continue
		}

		if command.Name() != ca.nameExpected {
			t.Errorf("Not equal")
		}
	}

}

func Test_Cli(t *testing.T) {
	dir := t.TempDir()

	mainFilePath := filepath.Join(dir, "cmd", "ox")
	if err := os.MkdirAll(mainFilePath, os.ModePerm); err != nil {
		t.Errorf("creating %s folder should not be error, but got %v", mainFilePath, err)
	}

	files := []struct {
		path    string
		name    string
		content string
	}{
		{
			path:    dir,
			name:    "go.mod",
			content: "module my/project",
		},
		{
			path:    mainFilePath,
			name:    "main.go",
			content: "package main",
		},
	}

	for _, f := range files {
		err := os.Chdir(f.path)
		if err != nil {
			t.Fatal(err)
		}

		err = os.WriteFile(f.name, []byte(f.content), 0444)
		if err != nil {
			t.Fatal(err)
		}
	}

	c := &cli{}

	t.Run("WrapOk running ox command in project dir with cmd/ox/main.go and go.mod", func(t *testing.T) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatal(err)
		}

		err = c.Wrap(context.Background(), []string{"ox"})
		if err != nil {
			t.Fatal(err)
		}
	})
}
