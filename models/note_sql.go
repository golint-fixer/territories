package models

import (
	// "database/sql"

	// "github.com/jmoiron/sqlx"
	"github.com/jinzhu/gorm"
)

type NoteSQL struct {
	// DB *sqlx.DB
	DB *gorm.DB
}

func (s *NoteSQL) Save(n *Note, userID uint, contactID uint) error {
	// var err error
	n.ContactID = contactID
	n.UserID = userID
	if n.ID == 0 {
		s.DB.Create(n)
		// 	if s.DB.DriverName() == "postgres" {
		// 		var result *sqlx.Rows
		// 		result, err = s.DB.NamedQuery("INSERT INTO notes (content, date, user_id, contact_id) VALUES (:content, :date, :user_id, :contact_id) RETURNING id", n)
		// 		result.Scan(&n.ID)
		// 	} else {
		// 		var result sql.Result
		// 		result, err = s.DB.NamedExec("INSERT INTO notes (content, date, user_id, contact_id) VALUES (:content, :date, :user_id, :contact_id)", n)
		// 		var id int64
		// 		id, err = result.LastInsertId()
		// 		n.ID = uint(id)
		// 	}
		// 	return err
		return s.DB.Error
	}

	// _, err = s.DB.NamedExec("UPDATE notes SET content=:content, date=:date, user_id=:user_id, contact_id=:contact_id WHERE id=:id", n)
	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contactID).Save(n)
	// return err
	return s.DB.Error
}

func (s *NoteSQL) Delete(n *Note, userID uint, contactID uint) error {
	n.ContactID = contactID
	n.UserID = userID

	// _, err := s.DB.NamedExec("DELETE FROM notes WHERE id=:id AND contact_id=:contact_id AND user_id=:user_id", n)
	// return err
	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contactID).Delete(n)
	return s.DB.Error
}

func (s *NoteSQL) FindByContact(contact Contact, userID uint) ([]Note, error) {
	var notes []Note
	s.DB.Where("user_id = ?", userID).Where("contact_id = ?", contact.ID).Find(&notes)
	// var err = s.DB.Select(&notes, "SELECT id, content, date FROM notes WHERE contact_id=? AND user_id=? ORDER BY date DESC", contact.ID, userID)
	// if err == sql.ErrNoRows || notes == nil {
	if s.DB.Error != nil {
		return make([]Note, 0), nil
	}
	// return notes, err
	return notes, s.DB.Error
}
