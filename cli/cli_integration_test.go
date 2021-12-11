//go:build integration
// +build integration

package cli_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/cli"
)

func TestNewApp(t *testing.T) {
	dir := t.TempDir()
	os.Chdir(dir)

	err := cli.Run(context.Background(), []string{"ox", "new", "coke"})
	if err != nil {
		t.Fatalf("error running new command: %v", err)
	}

	files := [][]string{
		{dir, "coke"},
		{dir, "coke", "go.mod"},
		{dir, "coke", "coke.go"},
	}

	for _, f := range files {
		file := filepath.Join(f...)
		if _, err := os.Stat(file); err == nil {
			continue
		}

		t.Fatalf("did not find: %v", file)
	}
}

func TestFixCommand(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting wd")
	}

	t.Cleanup(func() {
		os.Chdir(wd)
	})

	err = os.Chdir(t.TempDir())
	if err != nil {
		t.Fatalf("error moving to tempdir")
	}

	cmd := exec.Command("go", "install", "github.com/wawandco/ox/cmd/ox@v0.11.4")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("error running go install: %v", err)
	}

	cmd = exec.Command("ox", "new", "cocacola")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("error running ox new: %v", err)
	}

	err = os.Chdir("cocacola")
	if err != nil {
		t.Fatalf("error moving to the correct directory: %v", err)
	}

	err = cli.Run(context.Background(), []string{"ox", "fix"})
	if err != nil {
		t.Fatalf("error fixing app: %v", err)
	}

	files := [][]string{
		{dir, "coke"},
		{dir, "coke", "go.mod"},
		{dir, "coke", "coke.go"},
	}

	for _, f := range files {
		file := filepath.Join(f...)
		if _, err := os.Stat(file); err == nil {
			continue
		}

		t.Fatalf("did not find: %v", file)
	}

}
