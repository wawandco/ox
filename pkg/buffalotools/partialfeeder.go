package buffalotools

import (
	"io"
	"io/fs"
	"strings"
)

// NewPartialFeeder returns a partialFeeder that looks up for
// template files in the given FS and returns its contents as string.
func NewPartialFeeder(fs fs.FS) func(string) (string, error) {
	return func(name string) (string, error) {
		name = strings.TrimPrefix(name, "/")

		f, err := fs.Open(name)
		if err != nil {
			return "", err
		}

		b, err := io.ReadAll(f)
		if err != nil {
			return "", err
		}

		return string(b), nil
	}
}
