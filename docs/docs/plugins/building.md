---
title: "Building your CLI"
sidebar_position: 2
draft: true
---

One important part to mention is that the build command will not attempt to build the CLI folder. Instead the developer will need to do it when needed, either on the CLI

Either on your Dockerfile or your build system you should include something like `go build ./cmd/cli` to ensure that the cli binary gets to your running environment.