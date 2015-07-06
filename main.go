package main

import (
	"fmt"
	"net/http"
	"runtime"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/iogo-framework/application"
	"github.com/iogo-framework/cmd"
	"github.com/iogo-framework/logs"
	"github.com/iogo-framework/router"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
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
		cli.StringFlag{Name: "listen-host", Value: "0.0.0.0", Usage: "server listening host", EnvVar: "LISTEN_HOST"},
		cli.IntFlag{Name: "listen-port", Value: 8080, Usage: "server listening port", EnvVar: "LISTEN_PORT"},

		cli.StringFlag{Name: "sql-dialect", Value: "sqlite3", Usage: "database dialect ('sqlite' or 'postgres')", EnvVar: "SQL_DIALECT"},

		cli.StringFlag{Name: "postgres-host", Value: "postgres", Usage: "postgresql host", EnvVar: "POSTGRES_HOST"},
		cli.IntFlag{Name: "postgres-port", Value: 5432, Usage: "postgresql port", EnvVar: "POSTGRES_PORT"},
		cli.StringFlag{Name: "postgres-user", Value: "postgres", Usage: "postgresql user", EnvVar: "POSTGRES_USER"},
		cli.StringFlag{Name: "postgres-password", Value: "postgres", Usage: "postgresql password", EnvVar: "POSTGRES_PASSWORD"},
		cli.StringFlag{Name: "postgres-db", Value: "postgres", Usage: "postgresql database", EnvVar: "POSTGRES_DB"},

		cli.StringFlag{Name: "sqlite-path", Value: "/tmp/db.sqlite", Usage: "sqlite path", EnvVar: "SQLITE_PATH"},

		cli.BoolFlag{Name: "migrate, m", Usage: "migrate the database", EnvVar: "MIGRATE"},
		cli.BoolFlag{Name: "debug, d", Usage: "print debug information", EnvVar: "DEBUG"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

// TODO: Add Ping to the database when not migrating
func serve(ctx *cli.Context) error {
	var app *application.Application
	var err error

	if ctx.Bool("debug") {
		logs.Level(logs.DebugLevel)
	}

	var dialect, args string

	switch ctx.String("sql-dialect") {
	case "postgres":
		dialect = "postgres"
		args = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=verify-full",
			ctx.String("postgres-user"),
			ctx.String("postgres-password"),
			ctx.String("postgres-host"),
			ctx.Int("postgres-port"),
			ctx.String("postgres-db"),
		)
		logs.Debug("Database %s at %s", dialect, ctx.String("postgres-host"))
	case "sqlite3":
		fallthrough
	default:
		dialect = "sqlite3"
		args = ctx.String("sqlite-path")
		logs.Debug("Loading database %s at %s", dialect, args)
	}

	if ctx.Bool("migrate") {
		migrate(dialect, args)
		logs.Debug("Database migrated successfully")
	}

	app = application.New()
	if app.Components["DB"], err = sqlx.Connect(dialect, args); err != nil {
		return err
	}

	app.Components["Templates"] = make(map[string]*template.Template)
	app.Components["Mux"] = router.New()

	app.Use(router.Logger)
	app.Use(app.Apply)
	app.Use(cors)

	app.Post("/contacts", controllers.CreateContact)
	app.Options("/contacts", controllers.ContactCollectionOptions) // Required for CORS
	app.Get("/contacts", controllers.RetrieveContactCollection)

	app.Get("/contacts/:id", controllers.RetrieveContactByID)
	app.Patch("/contacts/:id", controllers.UpdateContactByID)
	app.Options("/contacts/:id", controllers.ContactOptions) // Required for CORS
	app.Delete("/contacts/:id", controllers.DeleteContactByID)

	return app.Serve(fmt.Sprintf("%s:%d", ctx.String("listen-host"), ctx.Int("listen-port")))
}

func cors(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func migrate(dialect string, args string) {
	db, err := gorm.Open(dialect, args)
	if err != nil {
		logs.Error(err)
		return
	}

	err = db.DB().Ping()
	if err != nil {
		logs.Error(err)
		return
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(false)

	db.AutoMigrate(models.Models()...)
}
