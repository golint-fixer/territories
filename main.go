package main

import (
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/Quorumsco/contact/controllers"
	"github.com/codegangsta/cli"
	"github.com/iogo-framework/application"
	"github.com/iogo-framework/cmd"
	"github.com/iogo-framework/router"
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
		cli.IntFlag{Name: "port, p", Value: 8080, Usage: "server listening port"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

func serve(ctx *cli.Context) error {
	var app *application.Application
	var err error

	if ctx.Bool("migrate") {
		//if err = databases.Migrate(models.Models()); err != nil {
		//return err
		//}
	}

	if app, err = application.New(); err != nil {
		return err
	}

	app.Mux = router.New()

	app.Use(router.Logger)
	app.Use(app.Apply)

	app.Post("/contacts", controllers.CreateContact)
	app.Options("/contacts", controllers.CreateContactOptions) // Required for CORS
	app.Get("/contacts", controllers.RetrieveContactCollection)

	app.Get("/contacts/:id", controllers.RetrieveContactByID)
	app.Patch("/contacts/:id", controllers.UpdateContactByID)
	app.Delete("/contacts/:id", controllers.DeleteContactByID)

	wd, _ := filepath.Abs("public")
	app.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir(wd))).ServeHTTP)

	app.Serve(ctx.Int("port"))

	return nil
}
