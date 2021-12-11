package npm

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/plugins/base/new"
)

type AfterInitializer struct{}

func (ai AfterInitializer) Name() string {
	return "npm/afterinitializer"
}

func (ai AfterInitializer) AfterInitialize(ctx context.Context, options new.Options) error {
	c := exec.CommandContext(ctx, "npm", "i", "--no-progress")
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}
