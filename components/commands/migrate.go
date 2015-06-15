package commands

import (
	"github.com/codegangsta/cli"

	"github.com/Quorumsco/contact/components/database"
	"github.com/Quorumsco/contact/models"
)

var CmdMigrate = cli.Command{
	Name:        "migrate",
	Usage:       "Migrate database",
	Description: `Automatic database migrations`,
	Before:      runMigrate,
	Action:      func(ctx *cli.Context) {},
	Flags:       []cli.Flag{},
}

func runMigrate(ctx *cli.Context) error {
	return database.Migrate(models.GetModels())
}
