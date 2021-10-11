---
title: "Writing a plugin"
sidebar_position: 2
---

Now that you understand the architecture and basic of plugins you can start writing your own plugins. This document will explain how to write a plugin and how to add it to the ox CLI. We will write a Command plugin, `greet` that will greet the user.

### Writing the plugin

Custom plugins are typically stored within the `/internal/ox/plugins` folder in your app, but feel free to put then anywhere else you consider. The plugin we will write is called `greet`, then the first thing we will do is implement the name method to return `greet`.

```go
// in /internal/ox/plugins/greet/greet.go
type GreetPlugin struct {}
func (p *GreetPlugin) Name() string {
    return `greet`
}
```

With that done we have only defined its an Ox plugin but we now need to make it a Command. The command interface looks like:

```go
type Command interface {
	Plugin              // Name() has already been implemented
	ParentName() string
	Run(context.Context, string, []string) error
}
```

So we need to implmement `Run` and `ParentName` in our GreetPlugin. Let's do it.

```go
// in internal/ox/plugins/greet/plugin.go
import (
    "context"
    "fmt"    
)

type Plugin struct {}
func (p Plugin) Name() string {
    return `greet`
}

func (p *Plugin) ParentName() string {
    return `` // We return an empty string since this is a root command.
}

func (p *Plugin) Run(context.Context, string, []string) error {
    // In here we just print out a message to the user.
    fmt.Println(`Hello! User`)
    return nil
}
```

And there we have our first plugin. 

### Connecting our plugin

One important part of the process is connecting the plugin to the CLI. To do this we add the plugin to the registry through the `Use` method of the CLI API. This all happens in the `cmd/ox/main.go`.

```go
// in cmd/ox/main.go
package main

import (
	"context"
	"fmt"
	"os"

	"myapp"
	_ "myapp/app/tasks"
	_ "myapp/app/models"
    
    // ** The plugin package is imported here so that it can be used. **
    "myapp/internal/ox/plugins/greet"

	"github.com/wawandco/ox/cli"
	"github.com/wawandco/ox/plugins/tools/soda"
)

// main function for the tooling cli, will be invoked by Ox
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its lifecycle.
func main() {
	cli.Use(soda.Plugins(myapp.Migrations)...) 
    // ...
    // ** The plugin is added here into the registry. **
    cli.Use(&greet.Plugin{})
    
    // Here we run the CLI.
	err := cli.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("[error] %v \n", err.Error())

		os.Exit(1)
	}
}
```

And thats it, once we invoke ox help we should see our new command in the list of commands.