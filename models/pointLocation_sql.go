// Definition of the structures and SQL interaction functions
package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// NoteSQL contains a Gorm client and the note and gorm related methods
type PointLocationSQL struct {
	DB *gorm.DB
}

// Save inserts a new note into the database
func (s *PointLocationSQL) Save(n *PointLocation, args PointLocationArgs) error {
	if n == nil {
		return errors.New("save: PointLocation is nil")
	}

	n.GroupID = args.PointLocation.GroupID
	if n.ID == 0 {
		err := s.DB.Create(n).Error
		s.DB.Last(n)
		return err
	}

	return s.DB.Where("group_id = ?", args.PointLocation.GroupID).Save(n).Error
}

// Delete removes a note from the database
func (s *PointLocationSQL) Delete(n *PointLocation, args PointLocationArgs) error {
	if n == nil {
		return errors.New("delete: PointLocation is nil")
	}

	return s.DB.Where("group_id = ?", args.PointLocation.GroupID).Delete(n).Error
}

// First returns a note from the database usin it's ID
func (s *PointLocationSQL) First(args PointLocationArgs) (*PointLocation, error) {
	var n PointLocation

	if err := s.DB.Where(args.PointLocation).First(&n).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &n, nil
}

// Find returns all the notes containing a given groupID from the database
func (s *PointLocationSQL) Find(args PointLocationArgs) ([]PointLocation, error) {
	var polygon []PointLocation

	err := s.DB.Where("group_id = ?", args.PointLocation.GroupID).Where("territory_id = ?", args.TerritoryID).Find(&polygon).Error
	if err != nil {
		return nil, err
	}

	return polygon, nil
}
