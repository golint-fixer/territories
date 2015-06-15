package controllers

import (
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/silverwyrda/iogo"
)

func getEnv(r *http.Request) (*sqlx.DB, map[string]*template.Template) {
	return iogo.GetContext(r).Env["DB"].(*sqlx.DB), iogo.GetContext(r).Env["Templates"].(map[string]*template.Template)
}
