package models

import "github.com/jmoiron/sqlx"

type TagDS interface {
	// Save(*Note) error
	// Delete(*Note) error
	// First(*Note) error
	// Find() ([]Note, error)
}

func TagStore(db *sqlx.DB) TagDS {
	return &TagSQL{DB: db}
}
