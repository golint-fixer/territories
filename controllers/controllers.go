package controllers

import (
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func getEnv(r *http.Request) (*sqlx.DB, map[string]*template.Template) {
	return iogo.GetContext(r).Env["Application"].DB.(*sqlx.DB), iogo.GetContext(r).Env["Application"].Templates.(map[string]*template.Template)
}
