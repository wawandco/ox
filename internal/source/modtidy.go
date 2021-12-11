package source

import (
	"os"
	"os/exec"

	"github.com/wawandco/ox/internal/log"
)

func RunModTidy(root string) error {
	err := os.Chdir(root)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Infof("Running: %s", cmd.String())

	return cmd.Run()
}
