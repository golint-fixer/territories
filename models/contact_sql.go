package models

import "github.com/jinzhu/gorm"

type ContactSQL struct {
	DB *gorm.DB
}

func (s *ContactSQL) Save(c *Contact, args ContactArgs) error {
	c.GroupID = args.GroupID
	if c.ID == 0 {
		s.DB.Create(c)

		return s.DB.Error
	}

	s.DB.Where("group_id = ?", args.GroupID).Save(c)

	return s.DB.Error
}

func (s *ContactSQL) Delete(c *Contact, args ContactArgs) error {
	s.DB.Where("group_id = ?", args.GroupID).Delete(c)

	return s.DB.Error
}

func (s *ContactSQL) First(args ContactArgs) (*Contact, error) {
	var c Contact

	s.DB.Where("group_id = ?", args.GroupID).Find(&c)

	return &c, s.DB.Error
}

func (s *ContactSQL) Find(args ContactArgs) ([]Contact, error) {
	var contacts []Contact

	s.DB.Where("group_id = ?", args.GroupID).Find(&contacts)
	if s.DB.Error != nil {
		return nil, s.DB.Error
	}

	return contacts, nil
}
