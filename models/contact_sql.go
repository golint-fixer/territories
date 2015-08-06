package models

import "github.com/jinzhu/gorm"

type ContactSQL struct {
	DB *gorm.DB
}

func (s *ContactSQL) Save(c *Contact, userID uint) error {
	c.UserID = userID
	if c.ID == 0 {
		s.DB.Create(c)

		return s.DB.Error
	}

	s.DB.Where("user_id = ?", userID).Save(c)

	return s.DB.Error
}

func (s *ContactSQL) Delete(c *Contact, userID uint) error {
	s.DB.Where("user_id = ?", userID).Delete(c)

	return s.DB.Error
}

func (s *ContactSQL) First(c *Contact, userID uint) error {
	s.DB.Where("user_id = ?", userID).Find(c)

	return s.DB.Error
}

func (s *ContactSQL) Find(userID uint) ([]Contact, error) {
	var contacts []Contact

	s.DB.Where("user_id = ?", userID).Find(&contacts)
	if s.DB.Error != nil {
		return make([]Contact, 0), nil
	}
	return contacts, s.DB.Error
}

func (s *ContactSQL) FindNotes(c *Contact, userID uint) error {
	var noteStore = NoteStore(s.DB)
	var err error

	c.Notes, err = noteStore.FindByContact(*c, userID)

	return err
}
