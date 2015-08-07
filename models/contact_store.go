package models

import "github.com/jinzhu/gorm"

type ContactDS interface {
	Save(*Contact, uint) error
	Delete(*Contact, uint) error
	First(*Contact, uint) error
	Find(uint) ([]Contact, error)

	FindNotes(*Contact, uint) error
	FindTags(*Contact) error
}

func ContactStore(db *gorm.DB) ContactDS {
	return &ContactSQL{DB: db}
}
