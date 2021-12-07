package standard

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/plugins/base/new"
)

type ModTidy struct{}

func (gag ModTidy) Name() string {
	return "mod-tidy"
}

func (gag ModTidy) AfterInitialize(ctx context.Context, options new.Options) error {
	err := os.Chdir(options.Folder)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(context.Background(), "go", "mod", "tidy")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
