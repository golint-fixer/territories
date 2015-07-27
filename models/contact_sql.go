package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ContactSQL struct {
	DB *sqlx.DB
}

func (s *ContactSQL) Save(c *Contact) error {
	if c.ID == 0 {
		var err error
		if s.DB.DriverName() == "postgres" {
			var result *sqlx.Rows
			result, err = s.DB.Queryx("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile) RETURNING id", c)
			result.Scan(&c.ID)
		} else {
			var result sql.Result
			result, err = s.DB.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile)", c)
			var id int64
			id, err = result.LastInsertId()
			c.ID = uint(id)
		}
		if err != nil {
			return err
		}

		return err
	}

	_, err := s.DB.NamedExec("UPDATE contacts SET firstname=:firstname, surname=:surname, married_name=:married_name, gender=:gender, birthdate=:birthdate, mail=:mail, phone=:phone, mobile=:mobile WHERE id=:id", c)
	if err != nil {
		return err
	}

	return nil
}

func (s *ContactSQL) Delete(c *Contact) error {
	_, err := s.DB.NamedExec("DELETE FROM contacts WHERE id=:id", c)
	if err != nil {
		return err
	}
	return nil
}

func (s *ContactSQL) First(c *Contact) error {
	if err := s.DB.Get(c, s.DB.Rebind("SELECT id, firstname, surname, phone FROM contacts WHERE id=? LIMIT 1"), c.ID); err != nil {
		return err
	}
	return nil
}

func (s *ContactSQL) Find() ([]Contact, error) {
	var contacts = make([]Contact, 0)
	if err := s.DB.Select(&contacts, "SELECT id, firstname, surname, phone FROM contacts ORDER BY surname DESC"); err != nil {
		if err == sql.ErrNoRows {
			return contacts, nil
		}
		return nil, err
	}
	return contacts, nil
}

func (s *ContactSQL) FindNotes(c *Contact) error {
	var err error

	noteStore := NoteStore(s.DB)
	c.Notes, err = noteStore.FindByContact(*c)

	return err
}
