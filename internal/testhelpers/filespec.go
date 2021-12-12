package testhelpers

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type Condition string

var (
	ConditionExists      = Condition("exists")
	ConditionNotExists   = Condition("not exists")
	ConditionContains    = Condition("contains")
	ConditionNotContains = Condition("not contains")
)

type FileSpec struct {
	Path      string
	Condition Condition
	Content   []string
}

func (fs FileSpec) Check(t *testing.T) {
	t.Helper()

	switch fs.Condition {
	case ConditionExists:
		_, err := os.Stat(fs.Path)
		if err != nil {
			t.Fatalf("file %s does not exist", fs.Path)
		}
	case ConditionNotExists:
		_, err := os.Stat(fs.Path)
		if err == nil {
			t.Fatalf("file %s does not exist", fs.Path)
		}

	case ConditionContains:
		_, err := os.Stat(fs.Path)
		if err != nil {
			t.Fatalf("file %s does not exist", fs.Path)
			return
		}

		f, err := os.Open(fs.Path)
		if err != nil {
			t.Fatalf("failed to open file %s: %s", fs.Path, err)
			return
		}

		dat, err := io.ReadAll(f)
		if err != nil {
			t.Fatalf("failed to read file %s: %s", fs.Path, err)
			return
		}

		for _, v := range fs.Content {
			if bytes.Contains(dat, []byte(v)) {
				continue
			}

			t.Fatalf("file %s does not contain %s", fs.Path, v)
			return
		}
	case ConditionContains:
		_, err := os.Stat(fs.Path)
		if err != nil {
			t.Fatalf("file %s does not exist", fs.Path)
			return
		}

		f, err := os.Open(fs.Path)
		if err != nil {
			t.Fatalf("failed to open file %s: %s", fs.Path, err)
			return
		}

		dat, err := io.ReadAll(f)
		if err != nil {
			t.Fatalf("failed to read file %s: %s", fs.Path, err)
			return
		}

		for _, v := range fs.Content {
			if !bytes.Contains(dat, []byte(v)) {
				continue
			}

			t.Fatalf("file %s contains %s", fs.Path, v)
			return
		}
	}

}

type FileSpecs []FileSpec

func (fss FileSpecs) CheckAll(t *testing.T) {
	t.Helper()

	for _, fs := range fss {
		fs.Check(t)
	}
}
