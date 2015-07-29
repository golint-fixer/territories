package models

import "github.com/jmoiron/sqlx"

type ContactDS interface {
	Save(*Contact) error
	Delete(*Contact) error
	First(*Contact) error
	Find(uint) ([]Contact, error)

	FindNotes(*Contact) error
}

func ContactStore(db *sqlx.DB) ContactDS {
	return &ContactSQL{DB: db}
}
