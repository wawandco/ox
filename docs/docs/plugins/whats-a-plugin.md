---
title: "What's a Plugin?"
sidebar_position: 1
---

The first question to ask is "What's a plugin?". 

     A plugin is a piece of code that is used to extend or replace functionality of the Ox CLI.

Plugins are written in Go and in order to be treated as such, they must comply with the "Plugin" interface.

```go
type Plugin interface {
    Name() string
}
```
Any type that matches that interface can be passed through the Ox CLI API as a plugin, and will be treated as such. then depending on the plugin, the goal of the plugin there are other interfaces to implement. Some of the examples are:

- `Command`: A plugin that implements the `Command` interface can be used to add a command to the Ox CLI.
- `Generator`: A plugin that implements `Generator` interface can be added to the generators list.

The plugins package contains some of the interfaces that plugins can implement, some other are defined by base plugins.

