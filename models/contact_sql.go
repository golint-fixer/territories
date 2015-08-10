package models

import "github.com/jinzhu/gorm"

type ContactSQL struct {
	DB *gorm.DB
}

func (s *ContactSQL) Save(c *Contact, groupID uint) error {
	c.GroupID = groupID
	if c.ID == 0 {
		s.DB.Create(c)

		return s.DB.Error
	}

	s.DB.Where("group_id = ?", groupID).Save(c)

	return s.DB.Error
}

func (s *ContactSQL) Delete(c *Contact, groupID uint) error {
	s.DB.Where("group_id = ?", groupID).Delete(c)

	return s.DB.Error
}

func (s *ContactSQL) First(c *Contact, groupID uint) error {
	s.DB.Where("group_id = ?", groupID).Find(c)

	return s.DB.Error
}

func (s *ContactSQL) Find(groupID uint) ([]Contact, error) {
	var contacts []Contact

	s.DB.Where("group_id = ?", groupID).Find(&contacts)
	if s.DB.Error != nil {
		return make([]Contact, 0), nil
	}
	return contacts, s.DB.Error
}

func (s *ContactSQL) FindNotes(c *Contact, groupID uint) error {
	var noteStore = NoteStore(s.DB)
	var err error

	c.Notes, err = noteStore.FindByContact(*c, groupID)

	return err
}

func (s *ContactSQL) FindTags(c *Contact) error {
	var tagStore = TagStore(s.DB)
	var err error

	c.Tags, err = tagStore.FindTagsByContact(*c)

	return err
}
