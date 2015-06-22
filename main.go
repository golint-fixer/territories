package main

import (
	"runtime"

	"./components/application"
	"./components/cmd"
	"./components/database"
	"./components/settings"
	"./models"
	"./routers"
	"github.com/codegangsta/cli"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd := cmd.New()
	cmd.Name = "contact"
	cmd.Usage = "Quorums contact backend"
	cmd.Version = "0.0.1"
	cmd.Before = serve
	cmd.Flags = append(cmd.Flags, []cli.Flag{
		cli.StringFlag{Name: "cpu, cpuprofile", Usage: "cpu profiling"},
		cli.BoolFlag{Name: "m, migrate", Usage: "migrate the database"},
		cli.StringFlag{Name: "port, p", Value: "8080", Usage: "server listening port"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

func serve(ctx *cli.Context) error {
	var app *application.Application
	var err error

	if ctx.Bool("migrate") {
		if err = database.Migrate(models.GetModels()); err != nil {
			return err
		}
	}

	settings.Port = ctx.String("port")

	if app, err = application.New(); err != nil {
		return err
	}

	app.Load(routers.URLs)
	app.Serve()

	return nil
}
