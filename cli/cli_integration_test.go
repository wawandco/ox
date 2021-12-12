//go:build integration
// +build integration

package cli_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/cli"
	"github.com/wawandco/ox/internal/testhelpers"
)

func TestNewApp(t *testing.T) {
	t.Run("NewAppOk", func(tt *testing.T) {
		testhelpers.WithinTempDir(tt, func(t *testing.T, dir string) {
			err := cli.Run(context.Background(), []string{"ox", "new", "coke", "-f"})
			if err != nil {
				t.Fatalf("error running new command: %v", err)
			}

			specs := testhelpers.FileSpecs{
				{
					Path:      filepath.Join(dir, "coke"),
					Condition: testhelpers.ConditionExists,
				},
				{
					Path:      filepath.Join(dir, "coke", "go.mod"),
					Condition: testhelpers.ConditionExists,
				},
				{
					Path:      filepath.Join(dir, "coke", "go.sum"),
					Condition: testhelpers.ConditionExists,
				},
			}

			specs.CheckAll(t)
		})
	})

	t.Run("FolderExists", func(tt *testing.T) {
		testhelpers.WithinTempDir(tt, func(t *testing.T, dir string) {
			err := os.MkdirAll(filepath.Join(dir, "coke"), 0644)
			if err != nil {
				t.Fatalf("error creating folder: %v", err)
			}

			err = cli.Run(context.Background(), []string{"ox", "new", "coke"})
			if err == nil {
				t.Fatalf("expected error running new command with folder existing")
			}
		})
	})

}

func TestNoCommand(t *testing.T) {
	cmd := exec.Command("go", "install", "github.com/wawandco/ox/cmd/ox")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("error running go install: %v", err)
	}

	cmd = exec.Command("ox")
	bs, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("error running ox: %v", err)
	}

	exp := []byte("no command provided, please provide one")
	if !bytes.Contains(bs, exp) {
		t.Fatalf("%v does not contain expected output: %v", bs, exp)
	}
}
