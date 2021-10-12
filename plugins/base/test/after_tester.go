package test

import (
	"context"

	plugins "github.com/wawandco/ox/plugins/core"
)

// AfterTester is suited for things that need to run after the tests
// cleanup and organization things, maybe reporting or collecting metrics.
type AfterTester interface {
	plugins.Plugin
	RunAfterTest(context.Context, string, []string) error
}
