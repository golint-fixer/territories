package models

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ContactSQL struct {
	DB *sqlx.DB
}

func (s *ContactSQL) Save(c *Contact, userID uint) error {
	var err error
	// We need to create a new record
	c.UserID = userID
	if c.ID == 0 {
		if s.DB.DriverName() == "postgres" {
			var result *sqlx.Rows
			if result, err = s.DB.Queryx("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile, user_id) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile, :UserID?) RETURNING id", c); err != nil {
				return err
			}
			result.Scan(&c.ID)
		} else {
			var result sql.Result
			if result, err = s.DB.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile, user_id) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile, :UserID?)", c); err != nil {
				return err
			}
			fmt.Println(c)
			var id int64
			id, err = result.LastInsertId()
			c.ID = uint(id)
		}
		return err
	}

	// We need to update the record
	_, err = s.DB.NamedExec("UPDATE contacts SET firstname=:firstname, surname=:surname, married_name=:married_name, gender=:gender, birthdate=:birthdate, mail=:mail, phone=:phone, mobile=:mobile WHERE id=:id AND user_id=:UserID?", c)

	return err
}

func (s *ContactSQL) Delete(c *Contact, userID uint) error {
	_, err := s.DB.NamedExec("DELETE FROM contacts WHERE id=:id AND UserID=?", c)
	return err
}

func (s *ContactSQL) First(c *Contact, userID uint) error {
	err := s.DB.Get(c, s.DB.Rebind("SELECT id, firstname, surname, phone FROM contacts WHERE id=? AND user_id=? LIMIT 1"), c.ID, userID)
	return err
}

func (s *ContactSQL) Find(userID uint) ([]Contact, error) {
	var contacts []Contact
	err := s.DB.Select(&contacts, "SELECT id, firstname, surname, phone FROM contacts WHERE user_id=? ORDER BY surname DESC", userID)
	if err == sql.ErrNoRows || contacts == nil {
		return make([]Contact, 0), nil
	}
	return contacts, err
}

func (s *ContactSQL) FindNotes(c *Contact) error {
	var noteStore = NoteStore(s.DB)
	var err error

	c.Notes, err = noteStore.FindByContact(*c)

	return err
}
