package ox

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wawandco/ox/internal/info"
)

// RenderFixer
type RenderFixer struct{}

func (rf RenderFixer) Name() string {
	return "ox/fixer/embed"
}

func (rf RenderFixer) Fix(ctx context.Context, root string, args []string) error {
	folder := filepath.Join(root, "app", "render")
	err := filepath.Walk(folder, func(path string, ii os.FileInfo, _ error) error {

		if ii.IsDir() {
			return nil
		}

		if !strings.HasSuffix(filepath.Base(ii.Name()), ".go") {
			return nil
		}

		bc, err := ioutil.ReadFile(path)
		if err != nil {

			return err
		}

		name, err := info.BuildName()
		if err != nil {
			return err
		}

		cc := strings.ReplaceAll(string(bc), name+".Templates.FindString", "buffalotools.NewPartialFeeder("+name+".Templates)")
		cc = strings.ReplaceAll(string(cc), "base.Templates.FindString", "buffalotools.NewPartialFeeder(base.Templates)")
		cc = strings.ReplaceAll(string(cc), "TemplatesBox:", "TemplatesFS:")
		cc = strings.ReplaceAll(string(cc), "AssetsBox:", "AssetsFS:")
		err = ioutil.WriteFile(path, []byte(cc), 0644)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
