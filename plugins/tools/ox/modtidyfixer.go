package ox

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/internal/log"
)

type ModTidyFixer struct{}

func (ef ModTidyFixer) Name() string {
	return "ox/fixer/modtidy"
}

func (ef ModTidyFixer) Fix(context.Context, string, []string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Infof("Running: %s", cmd.String())

	return cmd.Run()
}
