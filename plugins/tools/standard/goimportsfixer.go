package standard

import (
	"bytes"
	"context"
	"os"
	"path/filepath"

	"golang.org/x/tools/imports"
)

// GoImportsFixer runs goimports for the given root directory.
type GoImportsFixer struct{}

func (ef GoImportsFixer) Name() string {
	return "ox/fixer/adjustimports"
}

func (ef GoImportsFixer) Fix(ctx context.Context, root string, args []string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		res, err := imports.Process(path, src, nil)
		if err != nil {
			return err
		}

		if bytes.Equal(src, res) {
			return nil
		}

		return os.WriteFile(path, res, 0644)
	})

	return err
}
