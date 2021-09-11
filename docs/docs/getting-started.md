---
sidebar_position: 1
name: getting-started
title: Getting Started
---

### Installing the CLI

In order to get started with the Ox CLI you should install. You can grab the binary from the github repository or you can install from source, which is the recommended way. To install from source you should run:

```sh
go install github.com/wawandco/ox/cmd/ox@latest
```

Once this completes you should have the Ox binary in your terminal. You can test it and see if it works by running:

```sh
ox help
```

You should see something like:

```sh
[info] Using wawandco/ox/cmd/ox 

Ox allows to build apps with ease

Usage:
  ox [command]

Commands:
Command      Alias
  help          h       prints help text for the commands registered
  build         b       builds a buffalo app from within the root folder of the project
  dev           d       calls NPM or yarn to start webpack watching the assetst
  db                    database operation commands
  test                  provides the structure for test commands to run and be organized
  fix                   adapts the source code to comply with newer versions of the CLI
  generate      g       Allows to invoke registered generator plugins
  new                   Generates a new app with registered plugins
  task                  Runs grifts tasks previously imported in the CLI
  version       v       returns the current version of Ox CLI
```

Which means you're all set to start building your first application.

### Generating a new app

With the Ox CLI you can generate a new app from scratch, this process will build a folder structure with all that you need to get started writing you web application with Go. To generate this new app you can run:

```sh
ox new coke # Where  `coke` is the name of the app.
```

Once this has completed you should have a new folder called `coke` in your current working directory. This folder will contain all the files and folders that you need to get started with your web application.

### Setting up the database

Once your app has been generated you need to do one more step to run it, which is setting up the DB. Ox comes up with the `database` command to help you with this. To run the database command you can run:

```sh
ox db create 
```

This command will create a database instance for the development environment, in our case called `coke_development` which our application will connect to while in development mode.

### Running the app
Once the db has been setup you can start your app by running:

```sh
ox dev
```

And then visit the app in your browser at [http://localhost:3000/](http://localhost:3000/).
