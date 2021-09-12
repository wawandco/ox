---
title: "Test"
date: 2021-09-06T14:48:13-05:00
draft: false
---

The `test` command is probably one of the most used commands in ox, it allows to run the tests of the project. This command facilitates all the steps needed to run tests, from setting up the database to running the tests. You can use it by running:

```sh
ox test
```

In that case it will run all the tests in your source, and will print the results. However, sometimes you may want to run specific tests, or even a specific test file. Below you can find some example of how to run specific tests:

```sh
ox test ./actions/...                # Runs only the tests in the actions folder
ox test ./actions/... --run TestSome # Runs Tests in the actions folder which match TestSome
```



