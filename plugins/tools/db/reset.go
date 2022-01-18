package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/gobuffalo/pop/v6"
	"github.com/spf13/pflag"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins/core"
)

type ResetCommand struct {
	connectionName string
	// Other plugins that will run when reset is invoked
	resetters []Resetter

	flags *pflag.FlagSet
}

func (d ResetCommand) Name() string {
	return "reset"
}

func (d ResetCommand) HelpText() string {
	return "resets database specified in GO_ENV or --conn"
}

func (d ResetCommand) ParentName() string {
	return "database"
}

func (d *ResetCommand) Run(ctx context.Context, root string, args []string) error {
	conn := pop.Connections[d.connectionName]
	if conn == nil {
		return ErrConnectionNotFound
	}

	dial := conn.Dialect
	if dial == nil {
		return errors.New("provided connection is not a Resetter")
	}

	err := dial.DropDB()
	if err != nil {
		log.Warnf("could not drop database: %v\n", err)
	}
	log.Info("Database dropped")

	err = dial.CreateDB()
	if err != nil {
		log.Errorf("could not create database: %v\n", err)
		return err
	}
	log.Info("Database created")

	for _, resetter := range d.resetters {
		err := resetter.Reset(ctx, conn)
		if err != nil {
			log.Errorf("could not run resetter: %v\n", err)
			return err
		}
	}

	return nil
}

// RunBeforeTests will be invoked to reset the test database before
// tests run.
func (d *ResetCommand) RunBeforeTest(ctx context.Context, root string, args []string) error {
	if len(pop.Connections) == 0 {
		err := pop.LoadConfigFile()
		if err != nil {
			return fmt.Errorf("error on reset.RunBeforeTest: %w", err)
		}
	}

	conn := pop.Connections["test"]
	if conn == nil {
		return ErrConnectionNotFound
	}

	resetter := conn.Dialect
	if resetter == nil {
		return errors.New("provided connection is not a Resetter")
	}

	err := resetter.DropDB()
	if err != nil {
		log.Warnf("could not drop database: %v\n", err)
	}

	return resetter.CreateDB()
}

func (d *ResetCommand) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.Usage = func() {}
	d.flags.StringVarP(&d.connectionName, "conn", "", "development", "the name of the connection to use")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *ResetCommand) Flags() *pflag.FlagSet {
	return d.flags
}

func (d *ResetCommand) Receive(pls []core.Plugin) {
	for _, plugin := range pls {
		ptool, ok := plugin.(Resetter)
		if !ok {
			continue
		}

		d.resetters = append(d.resetters, ptool)
	}
}

// Resetter is something that should be run on reset.
type Resetter interface {
	Reset(context.Context, *pop.Connection) error
}
