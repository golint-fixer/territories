// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// NoteDS implements the NoteSQL methods
type PointLocationDS interface {
	Save(*PointLocation, PointLocationArgs) error
	Delete(*PointLocation, PointLocationArgs) error
	First(PointLocationArgs) (*PointLocation, error)
	Find(PointLocationArgs) ([]PointLocation, error)
}

// Notestore returns a NoteDS implementing CRUD methods for the notes and containing a gorm client
func PointLocationStore(db *gorm.DB) PointLocationDS {
	return &PointLocationSQL{DB: db}
}
