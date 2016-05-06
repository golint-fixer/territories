// Definition of the structures and SQL interaction functions
package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// ContactSQL contains a Gorm client and the contact and gorm related methods
type TerritorySQL struct {
	DB *gorm.DB
}

// Save inserts a new contact into the database
func (s *TerritorySQL) Save(c *Territory, args TerritoryArgs) error {
	if c == nil {
		return errors.New("save: Territory is nil")
	}

	c.GroupID = args.Territory.GroupID
	if c.ID == 0 {
		err := s.DB.Create(c).Error
		s.DB.Last(c)
		return err
	}

	return s.DB.Where("group_id = ?", args.Territory.GroupID).Save(c).Error
}

// Delete removes a contact from the database
func (s *TerritorySQL) Delete(c *Territory, args TerritoryArgs) error {
	if c == nil {
		return errors.New("delete: Territory is nil")
	}

	return s.DB.Where("group_id = ?", args.Territory.GroupID).Delete(c).Error
}

// First returns a Territory from the database using his ID
func (s *TerritorySQL) First(args TerritoryArgs) (*Territory, error) {
	var c Territory

	if err := s.DB.Where(args.Territory).First(&c).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	/*
		if err := s.DB.Where(c.AddressID).First(&c.Address).Error; err != nil {
			if err == gorm.RecordNotFound {
				return nil, nil
			}
			return nil, err
		}*/

	return &c, nil
}

// Find returns all the contacts with a given groupID from the database
func (s *TerritorySQL) Find(args TerritoryArgs) ([]Territory, error) {
	var territories []Territory

	err := s.DB.Where("group_id = ?", args.Territory.GroupID).Find(&territories).Error
	if err != nil {
		return nil, err
	}

	return territories, nil
}
