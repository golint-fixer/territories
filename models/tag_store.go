package models

import "github.com/jinzhu/gorm"

type TagDS interface {
	SaveTag(*Tag, Contact) error
	DeleteTag(*Tag, Contact) error
	FindTagsByContact(Contact) ([]Tag, error)
	FindTagById(*Tag, Contact) error
}

func TagStore(db *gorm.DB) TagDS {
	return &TagSQL{DB: db}
}
