package cmd

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/plugins/base/new"
)

func TestInitializer(t *testing.T) {
	t.Run("CompleteArgs", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join(root, "myapp"), 0777)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}
		ctx := context.Background()
		options := new.Options{
			Name:   "myapp",
			Module: "oosss/myapp",
			Folder: filepath.Join(root, "myapp"),
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		bmodels, err := os.ReadFile(filepath.Join(root, "myapp", "cmd", "myapp", "main.go"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bmodels, []byte(`var server = &servers.Simple{`)) {
			t.Fatal("should contain server definition")
		}

		if !bytes.Contains(bmodels, []byte(`bapp.Serve(server)`)) {
			t.Fatal("should use the server to start the app.")
		}

	})
}
