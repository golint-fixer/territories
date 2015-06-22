package routers

import (
	"net/http"
	"path/filepath"

	"../components/application"
	"../controllers"
	"github.com/silverwyrda/iogo"
)

func URLs(app *application.Application) {
	app.Use(iogo.Logger)
	app.Use(app.ApplyTemplates)
	app.Use(app.ApplyDB)

	app.Post("/v1.0/contacts", controllers.CreateContact)
	app.Options("/v1.0/contacts", controllers.CreateContactOptions)
	app.Get("/v1.0/contacts", controllers.RetrieveContactCollection)

	app.Get("/v1.0/contacts/:id", controllers.RetrieveContactByID)
	app.Patch("/v1.0/contacts/:id", controllers.UpdateContactByID)
	app.Delete("/v1.0/contacts/:id", controllers.DeleteContactByID)

	wd, _ := filepath.Abs("public")
	app.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir(wd))).ServeHTTP)
}
