package ox

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/internal/runtime"
)

// type InstallDependenciesFixer struct{}

type InstallDependenciesFixer struct{}

func (ef InstallDependenciesFixer) Name() string {
	return "ox/fixer/install-dependencies"
}

func (ef InstallDependenciesFixer) Fix(context.Context, string, []string) error {
	deps := []string{
		"github.com/gobuffalo/buffalo@v0.18",
		"github.com/wawandco/ox@" + runtime.Version,
	}

	cmd := exec.Command("go", "get")
	cmd.Args = append(cmd.Args, deps...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Infof("Running: %s", cmd.String())

	return cmd.Run()
}
