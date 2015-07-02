package main

import (
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
		cli.BoolFlag{Name: "migrate, m", Usage: "migrate the database"},
		cli.StringFlag{Name: "listen, l", Value: "0.0.0.0:8080", Usage: "server listening host:port"},
		cli.BoolFlag{Name: "debug, d", Usage: "print debug information"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

func serve(ctx *cli.Context) error {
	var app *application.Application
	var err error

	if ctx.Bool("migrate") {
		migrate()
		logs.Debug("Database migrated")
	}

	app = application.New()
	if app.Components["DB"], err = sqlx.Connect("sqlite3", "/tmp/contacts.db"); err != nil {
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

	app.Serve(ctx.String("listen"))

	return nil
}

func cors(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func migrate() {
	db, err := gorm.Open("sqlite3", "/tmp/contacts.db")
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
