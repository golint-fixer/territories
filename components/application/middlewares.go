package application

import (
	"net/http"

	"github.com/silverwyrda/iogo"
)

func (application *Application) ApplyTemplates(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		iogo.GetContext(r).Env["Templates"] = application.Templates
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (application *Application) ApplyDB(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		iogo.GetContext(r).Env["DB"] = application.DB
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
