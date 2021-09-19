---
title: "Overview"
sidebar_position: 0
---

The ox plugin system allows you to extend the functionality of ox, this is useful for development teams that want to streamline their development teams and incorporate custom workflows into the ox CLI. Ox ships with a base of plugins that we've find useful on the way but developers are encouraged to leverage the plugin system to add their own.

Some of the things you can achieve by using the plugin system are:
- Add new top level commands to the ox cli, e.g. `ox docs` to generate documentation
- Customize a build step in ox, e.g. to run tests before a build
- Extend the ox cli to add custom generators.
- Change the tooling used in the build step.

All that by writing your plugins or incorporating existing plugins into ox, that's the power of the OX plugin system.