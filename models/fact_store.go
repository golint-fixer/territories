// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// FactDS implements the FactSQL methods
type FactDS interface {
	Save(*Fact, FactArgs) error
	Delete(*Fact, FactArgs) error
	First(FactArgs) (*Fact, error)
	Find(FactArgs) ([]Fact, error)
}

// Factstore returns a FactDS implementing CRUD methods for the facts and containing a gorm client
func FactStore(db *gorm.DB) FactDS {
	return &FactSQL{DB: db}
}
