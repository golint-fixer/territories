package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type NoteSQL struct {
	DB *sqlx.DB
}

func (s *NoteSQL) Save(n *Note) error {
	var err error

	if n.ID == 0 {
		if s.DB.DriverName() == "postgres" {
			var result *sqlx.Rows
			result, err = s.DB.NamedQuery("INSERT INTO notes (content, date, user_id) VALUES (:content, :date, :user_id) RETURNING id", n)
			result.Scan(&n.ID)
		} else {
			var result sql.Result
			result, err = s.DB.NamedExec("INSERT INTO notes (content, date, user_id) VALUES (:content, :date, :user_id)", n)
			var id int64
			id, err = result.LastInsertId()
			n.ID = uint(id)
		}

		return err
	}

	_, err = s.DB.NamedExec("UPDATE notes SET content=:content, date=:date, user_id=:user_id WHERE id=:id", n)

	return err
}

func (s *NoteSQL) Delete(n *Note) error {
	_, err := s.DB.NamedExec("DELETE FROM notes WHERE id=:id", n)
	return err
}

func (s *NoteSQL) FindByContact(contact Contact) ([]Note, error) {
	var notes = make([]Note, 0)
	if err := s.DB.Select(&notes, "SELECT id, content, date FROM notes WHERE contact_id=? ORDER BY date DESC", contact.ID); err != nil {
		return nil, err
	}
	return notes, nil
}
