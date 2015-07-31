package controllers

import (
	"net/http"

	// "github.com/jmoiron/sqlx"
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/application"
	"github.com/quorumsco/router"
)

// func getDB(r *http.Request) *sqlx.DB {
// 	return router.Context(r).Env["Application"].(*application.Application).Components["DB"].(*sqlx.DB)
func getDB(r *http.Request) *gorm.DB {
	return router.Context(r).Env["Application"].(*application.Application).Components["DB"].(*gorm.DB)
}

func getUID(r *http.Request) uint {
	return router.Context(r).Env["UserID"].(uint)
}

//get gorm
