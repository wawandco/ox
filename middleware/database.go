package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/markbates/errx"
)

var errNonSuccess = errors.New("non success status code")

// Database returns a database connection provider middleware
// that will add the db into the `tx` context variable, in case its a
// PUT/POST/DELETE/PATCH method, it will wrap the
// handler in a transaction, otherwise it will just
// add the connection in the context.
//
// Provided middleware can also consider a seccond read only replica
// database connection (rodb) that will be used on idempotent operations if present.
func Database(db, rodb *pop.Connection) buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {

			// Setting the context on the DB.
			dbc := db.WithContext(c)

			switch c.Request().Method {
			default:
				fmt.Println(">> Idempotent <<")
				// If the passed read only db (rodb) is different than nil it should use it.
				// As this block is for read only operations.
				if rodb != nil {
					dbc = rodb.WithContext(c)
				}

				c.Set("tx", dbc)

				return h(c)
			case http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch:
				fmt.Println(">> Non Idempotent <<")

				couldBeDBorYourErr := dbc.Transaction(func(tx *pop.Connection) error {
					// setup logging
					start := tx.Elapsed
					defer func() {
						finished := tx.Elapsed
						elapsed := time.Duration(finished - start)
						c.LogField("db", elapsed)
					}()

					// add the transaction to the context
					c.Set("tx", tx)

					// call the next handler; if it errors stop and return the error
					if yourError := h(c); yourError != nil {
						return yourError
					}

					// check the response status code. if the code is NOT 200..399
					// then it is considered "NOT SUCCESSFUL" and an error will be returned
					if res, ok := c.Response().(*buffalo.Response); ok {
						if res.Status < 200 || res.Status >= 400 {
							return errNonSuccess
						}
					}

					// as far was we can tell everything went well
					return nil
				})

				// couldBeDBorYourErr could be one of possible values:
				// * nil - everything went well, if so, return
				// * yourError - an error returned from your application, middleware, etc...
				// * a database error - this is returned if there were problems committing the transaction
				// * a errNonSuccess - this is returned if the response status code is not between 200..399
				if couldBeDBorYourErr != nil && errx.Unwrap(couldBeDBorYourErr) != errNonSuccess {
					return couldBeDBorYourErr
				}

				return nil
			}
		}
	}
}
