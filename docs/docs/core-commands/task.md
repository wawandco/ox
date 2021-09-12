---
title: "Task"
date: 2021-09-06T14:48:13-05:00
draft: false
type: command
---

The task command allows to list an invoke `Grift` tasks defined in the `app/tasks` folder. To know more about `tasks` take a look at the `Tasks` documentation in the fundamental sections. To list the tasks available invoke the `tasks` command:

```sh
$ ox task

[info] Using github.com/wawandco/ox/cmd/ox 

There are 0 grift tasks available on this app:

task-name       Full Command
---------       ------------
create-users    ox task create-users
calculate-tax   ox task calculate-tax

```

Which will list the tasks available on the current app. To invoke one of these tasks you can run:

```sh
$ ox task <task-name>
```

Which will call the task specified in `task-name`.