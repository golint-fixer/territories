package models

import "github.com/jinzhu/gorm"

type NoteDS interface {
	Save(*Note, NoteArgs) error
	Delete(*Note, NoteArgs) error
	First(NoteArgs) (*Note, error)
	Find(NoteArgs) ([]Note, error)
}

func NoteStore(db *gorm.DB) NoteDS {
	return &NoteSQL{DB: db}
}
