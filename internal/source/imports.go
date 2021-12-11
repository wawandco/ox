package source

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/tools/imports"
)

func RunImports(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		src, err := ioutil.ReadFile(path)
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

		return ioutil.WriteFile(path, res, 0644)
	})

	return err
}
