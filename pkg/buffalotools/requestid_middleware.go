package buffalotools

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
)

// NewRequestIDMiddleware returns a middleware that adds a request ID to the
// context. It will first attempt to look if the passed headerkey is set in
// the request headers. If it is not set, it will generate a new UUID and set
// it in the logs.
func NewRequestIDMiddleware(headerKey string) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			id := c.Request().Header.Get(headerKey)
			if id == "" {
				id = fmt.Sprintf("%s", c.Value("request_id"))
			}

			c.LogField("request_id", id)

			return next(c)
		}
	}
}
