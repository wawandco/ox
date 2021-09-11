---
title: "Generate"
date: 2021-09-06T14:48:29-05:00
draft: false
type: command
---

Generate holds the subcommands that generate components in Ox applications, these components could be models, migrations, actions, templates. This is one of the commands where the benefit of the plugin system can be mostly exploded by writing custom generators that satisfy the requirements of the development team using the codebase. The generate command serves as the base for all those generate subcommands, but it does not do anything on its own.

Usage: 
```sh
ox generate <subcommand> [<args>...]
```

Some usage examples could be:
```sh
ox generate model users id:uuid name:string 
ox generate migration create_users_table
ox generate action users/create
ox generate template users/profile
```


