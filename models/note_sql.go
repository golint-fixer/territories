package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type NoteSQL struct {
	DB *sqlx.DB
}

func (s *NoteSQL) Save(n *Note, userID uint, contactID uint) error {
	var err error
	n.ContactID = contactID
	n.UserID = userID
	if n.ID == 0 {
		if s.DB.DriverName() == "postgres" {
			var result *sqlx.Rows
			result, err = s.DB.NamedQuery("INSERT INTO notes (content, date, user_id, contact_id) VALUES (:content, :date, :user_id, :contact_id) RETURNING id", n)
			result.Scan(&n.ID)
		} else {
			var result sql.Result
			result, err = s.DB.NamedExec("INSERT INTO notes (content, date, user_id, contact_id) VALUES (:content, :date, :user_id, :contact_id)", n)
			var id int64
			id, err = result.LastInsertId()
			n.ID = uint(id)
		}
		return err
	}

	_, err = s.DB.NamedExec("UPDATE notes SET content=:content, date=:date, user_id=:user_id, contact_id=:contact_id WHERE id=:id", n)
	return err
}

func (s *NoteSQL) Delete(n *Note, userID uint, contactID uint) error {
	n.ContactID = contactID
	n.UserID = userID
	_, err := s.DB.NamedExec("DELETE FROM notes WHERE id=:id AND contact_id=:contact_id AND user_id=:user_id", n)
	return err
}

func (s *NoteSQL) FindByContact(contact Contact, userID uint, contactID uint) ([]Note, error) {
	var notes []Note
	n.ContactID = contactID
	n.UserID = userID
	var err = s.DB.Select(&notes, "SELECT id, content, date FROM notes WHERE contact_id=? AND user_id=?ORDER BY date DESC", contact.ID, userID)
	if err == sql.ErrNoRows || notes == nil {
		return make([]Note, 0), nil
	}
	return notes, err
}
