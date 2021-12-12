package testhelpers

import (
	"os"
	"testing"
)

// Runs the passed funtion withn a temporary directory.
func WithinTempDir(t *testing.T, fn func(t *testing.T, dir string)) {
	t.Helper()

	wd, _ := os.Getwd()
	defer func() {
		err := os.Chdir(wd)
		if err != nil {
			t.Fatalf("failed to change working directory to %s: %s", wd, err)
		}
	}()

	td := t.TempDir()
	err := os.Chdir(td)
	if err != nil {
		t.Fatalf("failed to change directory to %s: %s", td, err)
		return
	}

	fn(t, td)
}
