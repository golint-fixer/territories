package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ContactSQL struct {
	DB *sqlx.DB
}

func (s *ContactSQL) Save(c *Contact) error {
	var err error

	// We need to create a new record
	if c.ID == 0 {
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

		return err
	}

	// We need to update the record
	_, err = s.DB.NamedExec("UPDATE contacts SET firstname=:firstname, surname=:surname, married_name=:married_name, gender=:gender, birthdate=:birthdate, mail=:mail, phone=:phone, mobile=:mobile WHERE id=:id", c)

	return err
}

func (s *ContactSQL) Delete(c *Contact) error {
	_, err := s.DB.NamedExec("DELETE FROM contacts WHERE id=:id", c)
	return err
}

func (s *ContactSQL) First(c *Contact) error {
	err := s.DB.Get(c, s.DB.Rebind("SELECT id, firstname, surname, phone FROM contacts WHERE id=? LIMIT 1"), c.ID)
	return err
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
