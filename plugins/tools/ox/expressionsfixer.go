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

// expressionsFixer
type expressionsFixer struct {
	expressions map[string]string
}

func (rf expressionsFixer) Name() string {
	return "ox/fixer/expressionfixer"
}

func (rf expressionsFixer) Fix(ctx context.Context, root string, args []string) error {
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

		for l, r := range rf.expressions {
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

func NewExpressionsFixer(expressions map[string]string) *expressionsFixer {
	return &expressionsFixer{expressions: expressions}
}
