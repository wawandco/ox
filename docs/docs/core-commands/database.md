---
title: "Database"
date: 2021-09-06T14:48:03-05:00
draft: false
type: command
---

The database command (db) contains database operation commands. Things like migrate, create your database live under the database command, the command allows you to do operations with your database.

### Subcommands
The Database command does not take any action when invoked directly, however, it  contains a set of subcommands that can be used to perform database operations which we will describe in detail below.
#### Create 
Creates a database on the specified connection, defaults to use the `development` connection. Some examples for its usage are:

```sh
$ ox db create
$ ox db create --conn=development
```
#### Migrate
Runs migrations on the specified direction and connection, defaults to use the `development` connection and the `up` direction. Some examples for its usage:

```sh
$ ox db migrate
$ ox db migrate up
$ ox db migrate down --conn development
```

#### Reset
Reset drops and recreates the database on the specified connection, defaults to use the `development` connection. Some examples for its usage:

```sh
$ ox db reset
$ ox db reset conn development
```



