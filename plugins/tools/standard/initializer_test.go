package standard_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/wawandco/ox/internal/runtime"
	"github.com/wawandco/ox/plugins/base/new"
	"github.com/wawandco/ox/plugins/tools/standard"
)

func TestInitializer(t *testing.T) {
	ini := standard.Initializer{}

	if ini.Name() != "standard/initializer" {
		t.Errorf("Expected 'standard/initializer' got '%s'", ini.Name())
	}

	err := os.Chdir(t.TempDir())
	if err != nil {
		t.Fatalf("could not change to temp dir: %s", err)
	}

	err = ini.Initialize(context.Background(), new.Options{
		Name:   "test",
		Module: "test",
	})

	if err != nil {
		t.Fatalf("error running initializer: %s", err)
	}

	if _, err := os.Stat("go.mod"); err != nil {
		t.Fatalf("did not generate go.mod: %s", err)
	}

	content, err := ioutil.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("could not read go.mod: %s", err)
	}

	contents := []string{
		"module test",
		fmt.Sprintf("github.com/wawandco/ox %s", runtime.Version),
	}

	scnt := string(content)
	for _, v := range contents {
		if !strings.Contains(scnt, v) {
			t.Fatalf("'%s' does not contain '%s'", scnt, v)
		}
	}
}
