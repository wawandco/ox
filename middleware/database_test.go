//go:build sqlite
// +build sqlite

package middleware_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/pop/v6"
	"github.com/wawandco/ox/middleware"
)

var deets = []*pop.ConnectionDetails{
	{Dialect: "sqlite3", URL: "sqlite://rodb.db?_busy_timeout=5000&_fk=true"},
	{Dialect: "sqlite3", URL: "sqlite://tx.db?_busy_timeout=5000&_fk=true"},
}

func TestDatabase(t *testing.T) {
	conns := map[string]*pop.Connection{}
	for _, cc := range deets {
		conn, err := pop.NewConnection(cc)
		if err != nil {
			t.Fatal(err)
		}

		conns[conn.Dialect.Details().Database] = conn
	}

	for _, conn := range conns {
		err := pop.CreateDB(conn)
		if err != nil {
			t.Fatal(err)
		}

		err = conn.Open()
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Cleanup(func() {
		for _, v := range conns {
			v.Close()
			err := pop.DropDB(v)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Nil Connections", func(t *testing.T) {
		app := buffalo.New(buffalo.Options{})

		var conn *pop.Connection
		hnd := func(c buffalo.Context) error {
			conn = c.Value("tx").(*pop.Connection)

			return nil
		}

		app.Use(middleware.Database(nil, nil))
		app.GET("/", hnd)
		app.POST("/", hnd)

		tt := httptest.New(app)
		r := tt.HTML("/").Get()

		if r.Code != http.StatusInternalServerError {
			t.Errorf("expected %d, got %d", http.StatusInternalServerError, r.Code)
		}

		if conn != nil {
			t.Errorf("expected nil, got %v", conn)
		}
	})

	t.Run("With one connection", func(t *testing.T) {
		app := buffalo.New(buffalo.Options{})
		app.Middleware.Clear()
		app.Use(middleware.Database(conns["tx.db"], nil))

		var conn *pop.Connection
		hnd := func(c buffalo.Context) error {
			conn = c.Value("tx").(*pop.Connection)

			return nil
		}

		app.GET("/", hnd)
		app.POST("/", hnd)

		tt := httptest.New(app)
		r := tt.HTML("/").Get()

		if r.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, r.Code)
		}

		if "sqlite3" != conn.Dialect.Name() {
			t.Errorf("expected %s, got %s", "sqlite3", conn.Dialect)
		}

		r = tt.HTML("/").Post(nil)

		if r.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, r.Code)
		}

		if "sqlite3" != conn.Dialect.Name() {
			t.Errorf("expected %s, got %s", "sqlite3", conn.Dialect)
		}
	})

	t.Run("With rodb connection", func(t *testing.T) {

		app := buffalo.New(buffalo.Options{})
		app.Middleware.Clear()
		app.Use(middleware.Database(conns["tx.db"], conns["rodb.db"]))

		var conn *pop.Connection
		hnd := func(c buffalo.Context) error {
			conn = c.Value("tx").(*pop.Connection)

			return nil
		}

		app.GET("/", hnd)
		app.POST("/", hnd)

		tt := httptest.New(app)
		r := tt.HTML("/").Get()

		if r.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, r.Code)
		}

		if "sqlite3" != conn.Dialect.Name() {
			t.Errorf("expected %s, got %s", "sqlite3", conn.Dialect)
		}

		if !strings.Contains(conn.Dialect.URL(), "rodb.db") {
			t.Errorf("expected %s to contain rodb.db but it did not", conn.Dialect.URL())
		}

		r = tt.HTML("/").Post(nil)

		if r.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, r.Code)
		}

		if !strings.Contains(conn.Dialect.URL(), "tx.db") {
			t.Errorf("expected %s to contain tx.db but it did not", conn.Dialect.URL())
		}
	})

}
