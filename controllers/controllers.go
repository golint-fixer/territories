package controllers

import (
	"html/template"
	"net/http"

	"github.com/iogo-framework/application"
	"github.com/iogo-framework/router"
	"github.com/jmoiron/sqlx"
)

func getEnv(r *http.Request) (*sqlx.DB, map[string]*template.Template) {
	return router.GetContext(r).Env["Application"].(*application.Application).DB.(*sqlx.DB), router.GetContext(r).Env["Application"].(*application.Application).Templates
}
