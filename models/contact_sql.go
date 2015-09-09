package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type ContactSQL struct {
	DB *gorm.DB
}

func (s *ContactSQL) Save(c *Contact, args ContactArgs) error {
	if c == nil {
		return errors.New("save: contact is nil")
	}

	c.GroupID = args.Contact.GroupID
	if c.ID == 0 {
		err := s.DB.Create(c).Error
		s.DB.Last(c)
		return err
	}

	return s.DB.Where("group_id = ?", args.Contact.GroupID).Save(c).Error
}

func (s *ContactSQL) Delete(c *Contact, args ContactArgs) error {
	if c == nil {
		return errors.New("delete: contact is nil")
	}

	return s.DB.Where("group_id = ?", args.Contact.GroupID).Delete(c).Error
}

func (s *ContactSQL) First(args ContactArgs) (*Contact, error) {
	var c Contact

	if err := s.DB.Where(args.Contact).First(&c).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

func (s *ContactSQL) Find(args ContactArgs) ([]Contact, error) {
	var contacts []Contact

	err := s.DB.Where("group_id = ?", args.Contact.GroupID).Find(&contacts).Error
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *ContactSQL) FindByMission(m *Mission, args ContactArgs) ([]Contact, error) {
	var contacts []Contact
	err := s.DB.Model(m).Related(&contacts, "Contacts").Error

	return contacts, err
}
