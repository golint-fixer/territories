package models

import "github.com/jmoiron/sqlx"

type NoteDS interface {
	Save(*Note) error
	Delete(*Note) error
	// First(*Note) error
	// Find() ([]Note, error)

	FindByContact(Contact) ([]Note, error)
}

func NoteStore(db *sqlx.DB) NoteDS {
	return &NoteSQL{DB: db}
}
