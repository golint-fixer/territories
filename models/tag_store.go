package models

import "github.com/jinzhu/gorm"

type TagDS interface {
	Save(*Tag, TagArgs) error
	Delete(*Tag, TagArgs) error
	Find(TagArgs) ([]Tag, error)
	// FindTagById(*Tag, Contact) error
}

func TagStore(db *gorm.DB) TagDS {
	return &TagSQL{DB: db}
}
