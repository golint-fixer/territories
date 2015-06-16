package routers

import (
	"net/http"
	"path/filepath"

	"github.com/Quorumsco/contact/components/application"
	"github.com/Quorumsco/contact/components/logs"
	"github.com/Quorumsco/contact/components/settings"
	"github.com/Quorumsco/contact/controllers"
	"github.com/silverwyrda/iogo"
)

func Init(app *application.Application) error {
	mux := iogo.New()
	//mux.Use(iogo.Logger)
	mux.Use(app.ApplyTemplates)
	mux.Use(app.ApplyDB)

	mux.Get("/api/v1.0/contacts", controllers.ContactList)
	mux.Get("/api/v1.0/contacts/:id", controllers.ContactByID)
	mux.Post("/api/v1.0/contacts/new", controllers.ContactNew)
	mux.Options("/api/v1.0/contacts/new", controllers.ContactNewOptions)

	wd, _ := filepath.Abs("public")
	mux.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir(wd))).ServeHTTP)

	logs.LogInfo("Listening on http://localhost:" + settings.Port)
	logs.LogCritical("%s", http.ListenAndServe(":"+settings.Port, mux))

	return nil
}
