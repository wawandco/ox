---
title: "Base Plugins"
sidebar_position: 2 
---

We mentioned the plugin registry in the [Architecture](/docs/plugins/architecture) section, the registry holds all the plugins that are available in the CLI. One important thing about the registry is that it does not come blank, rather it comes with a set of predefined plugins which we call `base`.

Ox is built on top of libraries we consider to be stable and reliable, and the list of plugins we use is a combination of these libraries. Among those libraries we use are:

- Buffalo (Framework)
- Pop (ORM)
- Docker (Dockerfile generated)
- Git (As the source control)
- Node (For frontend related work)
- Refresh (To hot-reload while in development)
- Webpack (For assets related work)
- Yarn (For frontend related work)

While we only have there stable libraries and plugins we tend to continuously evaluate the libraries we use and add/remove plugins depending on the evolution of the underlying libraries. We see base plugins as a foundation for anyone to get started with ox. 

The full list of base plugins can be found [here](https://github.com/wawandco/ox/blob/da3802e39c839864827d693f0fa6c2339626b0cb/tools/tools.go#L44).