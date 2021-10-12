package build

import "github.com/wawandco/ox/plugins/core"

// AfterBuilder interface allows to identify the things
// that will run after the build process has ended, things
// like cleanup and reverting go here
type AfterBuilder interface {
	core.Plugin
	RunAfterBuild(string, []string) error
}
