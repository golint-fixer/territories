package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/iogo-framework/application"
	"github.com/iogo-framework/cmd"
	"github.com/iogo-framework/databases"
	"github.com/iogo-framework/logs"
	"github.com/iogo-framework/router"
	"github.com/iogo-framework/settings"
	"github.com/quorumsco/contacts/controllers"
	"github.com/quorumsco/contacts/models"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd := cmd.New()
	cmd.Name = "contacts"
	cmd.Usage = "quorums contacts backend"
	cmd.Version = "0.0.1"
	cmd.Before = serve
	cmd.Flags = append(cmd.Flags, []cli.Flag{
		cli.StringFlag{Name: "config, c", Usage: "configuration file", EnvVar: "CONFIG"},
		cli.BoolFlag{Name: "debug, d", Usage: "debug log level"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

func serve(ctx *cli.Context) error {
	var app *application.Application
	var err error

	config, err := settings.Parse(ctx.String("config"))
	if err != nil && ctx.String("config") != "" {
		logs.Error(err)
	}

	if ctx.Bool("debug") || config.Debug() {
		logs.Level(logs.DebugLevel)
	}

	dialect, args, err := config.SqlDB()
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	logs.Debug("database type: %s", dialect)

	app = application.New()
	if app.Components["DB"], err = databases.InitSQLX(dialect, args); err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	logs.Debug("connected to %s", args)

	if config.Migrate() {
		if err := migrate(dialect, args); err != nil {
			logs.Critical(err)
			os.Exit(1)
		}
		logs.Debug("database migrated successfully")
	}

	app.Components["Mux"] = router.New()

	if ctx.Bool("debug") || config.Debug() {
		app.Use(router.Logger)
	}

	app.Use(app.Apply)
	app.Use(cors)

	app.Post("/contacts", controllers.CreateContact)
	app.Options("/contacts", controllers.ContactCollectionOptions) // Required for CORS
	app.Get("/contacts", controllers.RetrieveContactCollection)

	app.Get("/contacts/:id", controllers.RetrieveContactByID)
	app.Patch("/contacts/:id", controllers.UpdateContactByID)
	app.Options("/contacts/:id", controllers.ContactOptions) // Required for CORS
	app.Delete("/contacts/:id", controllers.DeleteContactByID)

	server, err := config.Server()
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	return app.Serve(server.String())
}

func cors(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func migrate(dialect string, args string) error {
	db, err := databases.InitGORM(dialect, args)
	if err != nil {
		return err
	}

	db.LogMode(true)

	db.AutoMigrate(models.Models()...)

	return nil
}
