package models

import "time"

type Mission struct {
	ID   uint       `gorm:"primary_key" json:"id"`
	Date *time.Time `json:"date,omitempty"`

	GroupID uint `sql:"not null" db:"group_id" json:"-"`

	Contacts []Contact `json:"contacts,omitempty" gorm:"many2many:mission_contacts;"`
}
