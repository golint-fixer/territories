package models

import "time"

type Mission struct {
	ID   uint
	Date *time.Time `json:"date,omitempty"`

	GroupID uint `sql:"not null" db:"group_id" json:"-"`

	Contacts []Contact `json:"tags,omitempty" gorm:"many2many:mission_contacts;"`
}
