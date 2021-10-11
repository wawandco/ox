---
title: "Middleware"
date: 2021-09-02T14:45:49-05:00
draft: false
sidebar_position: 2
---

Middleware functions are key component of the Buffalo framework, in Buffalo a Middleware facilitates tasks that need to be done across multiple actions or the entire application. Examples for middleware use case include:

- Logging
- Authentication
- Authorization
- Caching
- Rate limiting
- Loading common data

And in general, middleware functions are used to perform tasks that need to be done before or after an action is executed. 

### Middleware Anatomy
Middleware function must have the [buffalo.Middleware](github.com/gobuffalo/buffalo) type which receives a handler and returns another one.

```go
type Middleware func(handler buffalo.Handler) buffalo.Handler
```

As an example you can see next a middleware function that adds a request ID to the context, so it can later be used by the logger for debugging purposes.

```go
// SetRequestID sets a unique request ID in the context,
// this is just an example but could be used for logging.
func SetRequestID(h buffalo.Handler) buffalo.Handler {
    // Returns the original handler wrapped within a Handler function
    // which logs something before calling the original handler.
    return func (c buffalo.Context) error {
        c.Set("requestID", uuid.NewV4().String())
        return h(c)
    }
}
```
To use a middleware Buffalo applications (and Groups) have the the `Use` method which takes one or more middleware functions and applies it to the application. 

### Ox default Middleware
Ox generated applications ship with default middleware for common things in Buffalo applications, these middleware are defined in the `middleware` folder and used in `app/routes.go`.

```go
// in routes.go
// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)
    ...
```

There are 3 default middleware functions that are used by the application. Lets explain what these are:

- middleware.Transaction: Takes care of setting up a `tx` field in the context for each request, so handlers can use it to access the database.
- middleware.ParameterLogger: Logs the parameters of the request.
- middleware.CSRF: Adds a CSRF token to the context for use in forms.


For more info on Middleware see [Middleware](#middleware).
