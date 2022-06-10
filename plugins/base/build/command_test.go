package build

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/wawandco/ox/plugins/tools/standard"
	"github.com/wawandco/ox/plugins/tools/yarn"
)

func TestSetEnv(t *testing.T) {
	b := &Command{}

	t.Run("Unset", func(t *testing.T) {
		err := b.setenv()
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		env := os.Getenv("GO_ENV")
		if env != "production" {
			t.Errorf("GO_ENV should have been production, got %v", env)
		}
	})

	t.Run("Set to development", func(t *testing.T) {
		err := os.Setenv("GO_ENV", "development")
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		err = b.setenv()
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		env := os.Getenv("GO_ENV")
		if env != "development" {
			t.Errorf("GO_ENV should have been %v, got %v", "development", env)
		}

	})
}

func TestCommand_Run(t *testing.T) {
	st := &standard.Builder{}
	yb := &yarn.Plugin{}

	b := &Command{
		builders:       []Builder{st},
		beforeBuilders: []BeforeBuilder{yb},
	}

	root := t.TempDir()
	err := os.Chdir(root)
	if err != nil {
		t.Fatalf("could not change to temp dir: %s", err)
	}

	t.Run("go mod missing error", func(t *testing.T) {
		err := b.Run(context.Background(), root, []string{"build"})

		if err.Error() != "open go.mod: no such file or directory" {
			t.Errorf("should return 'open go.mod: no such file or directory' error, got %v", err)
		}
	})

	file := `module test`
	err = ioutil.WriteFile("go.mod", []byte(file), 0444)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("build error", func(t *testing.T) {
		err := b.Run(context.Background(), root, []string{"build"})

		if err.Error() != "exit status 1" {
			t.Errorf("should return 'exit status 1' error, got %v", err)
		}
	})

}
