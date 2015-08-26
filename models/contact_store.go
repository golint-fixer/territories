package models

import "github.com/jinzhu/gorm"

type ContactDS interface {
	Save(*Contact, ContactArgs) error
	Delete(*Contact, ContactArgs) error
	First(ContactArgs) (*Contact, error)
	Find(ContactArgs) ([]Contact, error)
	FindByMission(*Mission, ContactArgs) ([]Contact, error)

	// FindNotes(*Contact, *ContactArgs) error
	// FindTags(*Contact) error
}

func ContactStore(db *gorm.DB) ContactDS {
	return &ContactSQL{DB: db}
}
