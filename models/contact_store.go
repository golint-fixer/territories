package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ContactDS interface {
	Save(*Contact) error
	Delete(*Contact) error
	First(*Contact) error
	Find() ([]Contact, error)
}

type ContactSQL struct {
	DB *sqlx.DB
}

func ContactStore(db *sqlx.DB) ContactDS {
	return &ContactSQL{DB: db}
}

func (s *ContactSQL) Save(c *Contact) error {
	if c.ID == 0 {
		var result sql.Result
		var err error
		if s.DB.DriverName() == "postgres" {
			result, err = s.DB.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile) RETURNING id", c)
		} else {
			result, err = s.DB.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile)", c)
		}
		if err != nil {
			return err
		}

		c.ID, err = result.LastInsertId()
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
	if err := s.DB.Get(c, s.DB.Rebind("SELECT * FROM contacts WHERE id=? LIMIT 1"), c.ID); err != nil {
		return err
	}
	return nil
}

func (s *ContactSQL) Find() ([]Contact, error) {
	contacts := make([]Contact, 0)
	if err := s.DB.Select(&contacts, "SELECT id, firstname, surname FROM contacts ORDER BY surname DESC"); err != nil {
		if err == sql.ErrNoRows {
			return contacts, nil
		}
		return nil, err
	}
	return contacts, nil
}
