package buffalotools

import (
	"io/fs"
	"log"

	"github.com/gobuffalo/pop/v6"
	"github.com/markbates/oncer"
)

var (
	// connections opened by the DatabaseProvider
	connections = map[string]*pop.Connection{}
)

// DatabaseProvider returns a function that returns the database connection
// for the current environment.
func DatabaseProvider(config fs.FS) func(name string) *pop.Connection {
	// Loading connections from database.yml in the pop.Connections
	// variable for later usage.
	bf, err := config.Open("database.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = pop.LoadFrom(bf)
	if err != nil {
		log.Fatal(err)
	}

	return func(name string) *pop.Connection {
		// Only open the connection once.
		oncer.Do("pop:connection:"+name, func() {
			c, err := pop.Connect(name)
			if err != nil {
				log.Fatal(err)
			}

			connections[name] = c
		})

		return connections[name]
	}
}
