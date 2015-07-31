package models

import (
	// "database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	// "github.com/jmoiron/sqlx"
)

type ContactSQL struct {
	//DB *sqlx.DB
	DB *gorm.DB
}

func (s *ContactSQL) Save(c *Contact, userID uint) error {
	// var err error
	// We need to create a new record
	c.UserID = userID
	if c.ID == 0 {
		s.DB.Create(c)
		// if s.DB.DriverName() == "postgres" {
		// 	var result *sqlx.Rows
		// 	if result, err = s.DB.Queryx("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile, user_id) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile, :user_id) RETURNING id", c); err != nil {
		// 		return err
		// 	}
		// 	result.Scan(&c.ID)
		// } else {
		// 	var result sql.Result
		// 	if result, err = s.DB.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile, user_id) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile, :user_id)", c); err != nil {
		// 		return err
		// 	}
		// 	fmt.Println(c)
		// 	var id int64
		// 	id, err = result.LastInsertId()
		// 	c.ID = uint(id)

		// }
		// return err
		return s.DB.Error
	}

	// We need to update the record
	// _, err = s.DB.NamedExec("UPDATE contacts SET firstname=:firstname, surname=:surname, married_name=:married_name, gender=:gender, birthdate=:birthdate, mail=:mail, phone=:phone, mobile=:mobile WHERE id=:id AND user_id=:user_id", c)
	s.DB.Where("user_id = ?", userID).Save(c)

	// return err
	return s.DB.Error
}

func (s *ContactSQL) Delete(c *Contact, userID uint) error {
	// _, err := s.DB.NamedExec("DELETE FROM contacts WHERE id=:id AND user_id=?", c)
	// return err
	s.DB.Where("user_id = ?", userID).Delete(c)
	return s.DB.Error
}

func (s *ContactSQL) First(c *Contact, userID uint) error {
	// err := s.DB.Get(c, s.DB.Rebind("SELECT id, firstname, surname, phone FROM contacts WHERE id=? AND user_id=? LIMIT 1"), c.ID, userID)
	s.DB.Where("user_id = ?", userID).Where("id = ?", c.ID).Find(c)
	fmt.Println(c)
	return s.DB.Error
	//par pointeur

}

func (s *ContactSQL) Find(userID uint) ([]Contact, error) {
	var contacts []Contact

	s.DB.Where("user_id = ?", userID).Find(&contacts)

	// err := s.DB.Select(&contacts, "SELECT id, firstname, surname, phone FROM contacts WHERE user_id=? ORDER BY surname DESC", userID)
	// if err == sql.ErrNoRows || contacts == nil {
	if s.DB.Error != nil {
		return make([]Contact, 0), nil
	}
	// return contacts, err
	return contacts, s.DB.Error
}

func (s *ContactSQL) FindNotes(c *Contact, userID uint) error {
	var noteStore = NoteStore(s.DB)
	var err error

	c.Notes, err = noteStore.FindByContact(*c, userID)

	return err
}
