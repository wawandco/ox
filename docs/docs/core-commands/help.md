---
title: "Help"
date: 2021-09-06T14:48:45-05:00
draft: false
type: command
short: ""
---

The help command provides valuable information for each of the commands in the CLI. The help command shows the information it has about a command or subcommand. You can see the top level help by running:

```
ox help
```

Also, you can get specific help for a particular command by running `ox help [command]`. For example:

```
ox help new
ox help database create # shoes the information for the `create` subcommand .
```

The help command produces output similar to the following:
```sh
$ ox help database
[info] Using github.com/wawandco/ox/cmd/ox 

database operation commands

Usage:
  ox db [subcommand]

Subcommands:
  create        creates database in GO_ENV or --conn flag
  drop          drops database in GO_ENV or --conn flag
  reset         resets database specified in GO_ENV or --conn
```