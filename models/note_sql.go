package models

import "github.com/jinzhu/gorm"

type NoteSQL struct {
	DB *gorm.DB
}

func (s *NoteSQL) Save(n *Note, userID uint, contactID uint) error {
	n.ContactID = contactID
	n.UserID = userID
	if n.ID == 0 {
		s.DB.Create(n)

		return s.DB.Error
	}

	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contactID).Save(n)

	return s.DB.Error
}

func (s *NoteSQL) Delete(n *Note, userID uint, contactID uint) error {
	n.ContactID = contactID
	n.UserID = userID

	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contactID).Delete(n)

	return s.DB.Error
}

func (s *NoteSQL) FindByContact(contact Contact, userID uint) ([]Note, error) {
	var notes []Note
	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contact.ID).Find(&notes)
	if s.DB.Error != nil {
		return make([]Note, 0), nil
	}

	return notes, s.DB.Error
}

func (s *NoteSQL) FindNoteById(n *Note, userID uint, noteID uint, contactID uint) error {
	n.ContactID = contactID
	n.ID = noteID
	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contactID).Find(n)

	return s.DB.Error
}
