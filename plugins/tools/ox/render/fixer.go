package render

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/wawandco/ox/internal/info"
)

// Fixer
type Fixer struct{}

func (rf Fixer) Name() string {
	return "render/fixer"
}

func (rf Fixer) Fix(ctx context.Context, root string, args []string) error {
	folder := filepath.Join(root, "app", "render")
	err := filepath.Walk(folder, func(path string, ii os.FileInfo, _ error) error {

		if ii.IsDir() {
			return nil
		}

		if !strings.HasSuffix(filepath.Base(ii.Name()), ".go") {
			return nil
		}

		bc, err := os.ReadFile(path)
		if err != nil {

			return err
		}

		name, err := info.BuildName()
		if err != nil {
			return err
		}

		cc := strings.ReplaceAll(string(bc), name+".Templates.FindString", "buffalotools.NewPartialFeeder(templates.FS())")
		cc = strings.ReplaceAll(string(cc), "base.Templates.FindString", "buffalotools.NewPartialFeeder(templates.FS())")
		cc = strings.ReplaceAll(string(cc), "TemplatesBox:", "TemplatesFS:")
		cc = strings.ReplaceAll(string(cc), "AssetsBox:", "AssetsFS:")
		cc = strings.ReplaceAll(string(cc), "base.Templates", "templates.FS()")
		cc = strings.ReplaceAll(string(cc), "base.Assets", "public.FS()")
		cc = strings.ReplaceAll(string(cc), name+".Templates", "templates.FS()")
		cc = strings.ReplaceAll(string(cc), name+".Assets", "public.FS()")
		err = os.WriteFile(path, []byte(cc), 0644)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
