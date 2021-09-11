---
title: "Build"
date: 2021-09-06T14:47:48-05:00
draft: false
type: command
---

The build command is in charge of building the final binary for the Ox app. It invokes things like the node build process before packing the binary embedding the asset files. This command attempts to build the Go application that's located within `cmd/[app-name]/main.go`. The end result of this command is a self-contained binary that can be executed to run the app. 

Invoking the build command can be done by running:

```sh
ox build
```

Some other examples that could come handy are:

```sh
ox build -o name            # changes the name of the built output
ox build --static -o name   # buils a static binary
ox build --tags netdns=go   # specifying build tags to the go build commands
```

To know more about the command you can always use the [`help`](/docs/core-commands/help) command.

### The build process

### Multiple binaries
One important thing to mention here is that Ox recommended approach is that an application will have multiple binaries built, one of each purpose. That way when a binary is invoked in production it will know the single task it will run.

On a typical app we could have:
```
yourapp
  app
  cmd
    yourapp
      main.go  // The binary that serves the app handlers and routes
    ox
      main.go  // a binary for CLI cron tasks, migrations etc
    worker
      main.go  // a worker binary for things like Temporal or Faktory.
```

And the Dockerfile could just build those to be ready in the Dockerfile.
