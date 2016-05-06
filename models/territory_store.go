// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// TerritoryDS implements the TerritorySQL methods
type TerritoryDS interface {
	Save(*Territory, TerritoryArgs) error
	Delete(*Territory, TerritoryArgs) error
	First(TerritoryArgs) (*Territory, error)
	Find(TerritoryArgs) ([]Territory, error)

	// FindNotes(*Contact, *ContactArgs) error
	// FindTags(*Contact) error
}

// Contactstore returns a ContactDS implementing CRUD methods for the contacts and containing a gorm client
func TerritoryStore(db *gorm.DB) TerritoryDS {
	return &TerritorySQL{DB: db}
}
