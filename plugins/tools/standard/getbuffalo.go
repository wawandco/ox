package standard

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/plugins/base/new"
)

type GetBuffalo struct{}

func (gag GetBuffalo) Name() string {
	return "getbuffalo"
}

// Getting correct Buffalo version.
func (gag GetBuffalo) AfterInitialize(ctx context.Context, options new.Options) error {
	err := os.Chdir(options.Folder)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(context.Background(), "go", "get", "github.com/gobuffalo/buffalo@v0.18")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
