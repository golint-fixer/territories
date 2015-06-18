package models

import (
	"database/sql"

	"github.com/Quorumsco/contact/components/database"
)

type ContactStore interface {
	Save(*Contact) error
	Delete(*Contact) error
	First(*Contact) error
	Find() ([]Contact, error)
}

type ContactSQL struct {
	DB *database.DB
}

func NewContactStore(db *database.DB) ContactStore {
	return &ContactSQL{DB: db}
}

func (s *ContactSQL) Save(c *Contact) error {
	if c.ID == 0 {
		result, err := s.DB.SQLX.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile)", c)
		if err != nil {
			return err
		}

		c.ID, err = result.LastInsertId()
		return err
	}

	_, err := s.DB.SQLX.NamedExec("UPDATE contacts SET firstname=:firstname, surname=:surname, married_name=:married_name, gender=:gender, birthdate=:birthdate, mail=:mail, phone=:phone, mobile=:mobile WHERE id=:id", c)
	if err != nil {
		return err
	}

	return nil
}

func (s *ContactSQL) Delete(c *Contact) error {
	_, err := s.DB.SQLX.NamedExec("DELETE FROM contacts WHERE id=:id", c)
	if err != nil {
		return err
	}
	return nil
}

func (s *ContactSQL) First(c *Contact) error {
	if err := s.DB.SQLX.Get(c, "SELECT * FROM contacts WHERE id=? LIMIT 1", c.ID); err != nil {
		return err
	}
	return nil
}

func (s *ContactSQL) Find() ([]Contact, error) {
	contacts := make([]Contact, 0)
	if err := s.DB.SQLX.Select(&contacts, "SELECT id, firstname, surname FROM contacts ORDER BY surname DESC"); err != nil {
		if err == sql.ErrNoRows {
			return contacts, nil
		}
		return nil, err
	}
	return contacts, nil
}
