package models

import "github.com/jmoiron/sqlx"

type ContactDS interface {
	Save(*Contact, uint) error
	Delete(*Contact, uint) error
	First(*Contact, uint) error
	Find(uint) ([]Contact, error)

	FindNotes(*Contact, uint) error
}

func ContactStore(db *sqlx.DB) ContactDS {
	return &ContactSQL{DB: db}
}
