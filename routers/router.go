package routers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Quorumsco/contact/components/application"
	"github.com/Quorumsco/contact/controllers"
	"github.com/silverwyrda/iogo"
)

// Since the app is fairly simple, no need to separate
// all the urls
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

	log.Fatal(http.ListenAndServe(":8080", mux))

	return nil
}
