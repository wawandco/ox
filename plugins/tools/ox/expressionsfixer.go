package ox

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wawandco/ox/internal/info"
)

var exps = map[string]string{
	"middleware.Database(models.DB":        "buffalotools.DatabaseMiddleware(models.DB",
	"github.com/wawandco/ox/middleware":    "github.com/wawandco/ox/pkg/buffalotools",
	".ServeFiles(\"/\", {{.Name}}.Assets)": ".ServeFiles(\"/\", http.FS(public.FS()))",
	".ServeFiles(\"/\", base.Assets)":      ".ServeFiles(\"/\", http.FS(public.FS()))",
}

// ExpressionsFixer for buffalo/ox expressions that may have changed.
type ExpressionsFixer struct{}

func (rf ExpressionsFixer) Name() string {
	return "buffalo/expressionsfixer"
}

func (rf ExpressionsFixer) Fix(ctx context.Context, root string, args []string) error {
	err := filepath.Walk(root, func(path string, ii os.FileInfo, _ error) error {
		if ii.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		bn, err := info.BuildName()
		if err != nil {
			return err
		}

		cc, err := ioutil.ReadFile(path)
		if err != nil {

			return err
		}

		for l, r := range exps {
			l = strings.Replace(l, "{{.Name}}", bn, -1)
			r = strings.Replace(r, "{{.Name}}", bn, -1)

			cc = bytes.ReplaceAll(cc, []byte(l), []byte(r))

			err = ioutil.WriteFile(path, cc, 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
