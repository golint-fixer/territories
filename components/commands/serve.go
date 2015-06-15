package commands

import (
	"github.com/codegangsta/cli"

	"github.com/Quorumsco/contact/components/application"
	"github.com/Quorumsco/contact/components/settings"
	"github.com/Quorumsco/contact/routers"
)

var CmdServe = cli.Command{
	Name:  "serve",
	Usage: "Start quorums web serve",
	Description: `This is the only thing you need to run,
    and it takes care of all the other things for you`,
	Before: runServe,
	Action: func(ctx *cli.Context) {},
	Flags: []cli.Flag{
		cli.StringFlag{"port, p", "8080", "Server listening port", ""},
		cli.BoolFlag{"migrate, m", "Migrate before starting", ""},
	},
}

func runServe(ctx *cli.Context) error {
	if ctx.Bool("migrate") {
		if err := runMigrate(ctx); err != nil {
			return err
		}
	}

	var app = &application.Application{}
	err := app.Init()
	if err != nil {
		return err
	}

	settings.Port = ctx.String("port")

	err = routers.Init(app)
	if err != nil {
		return err
	}

	return nil
}
