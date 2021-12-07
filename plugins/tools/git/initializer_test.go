package git

import (
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

		keeps := []string{
			"migrations",
			"public",
		}

		for _, k := range keeps {
			_, err := os.Stat(filepath.Join(root, "myapp", k, ".gitkeep"))
			if err != nil {
				t.Fatal("should have created the file")
			}
		}
	})

	t.Run("With_Git_Ignore", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join(root, "anotherapp"), 0777)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}
		ctx := context.Background()
		options := new.Options{
			Name:   "anotherapp",
			Module: "oosss/anotherapp",
			Folder: filepath.Join(root, "anotherapp"),
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		rootGitIgnore := filepath.Join(options.Folder, ".gitignore")
		_, err = os.Stat(rootGitIgnore)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create .gitignore file , %v", err)
		}
	})
}
