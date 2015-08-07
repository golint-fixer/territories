package models

import "github.com/jinzhu/gorm"

type NoteDS interface {
	Save(*Note, uint, uint) error
	Delete(*Note, uint, uint) error

	FindByContact(Contact, uint) ([]Note, error)
	FindById(*Note, uint, uint, uint) error
}

func NoteStore(db *gorm.DB) NoteDS {
	return &NoteSQL{DB: db}
}
