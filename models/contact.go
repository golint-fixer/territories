package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Position struct {
	//ID float64

	X float64
	Y float64

	Latitude  float64
	Longitude float64
}

type Address struct {
	ID int64

	HouseNumber string
	Street      string
	PostalCode  string
	City        string
	County      string // Département
	State       string // Région
	Country     string
	Addition    string // Complément d'adresse

	PollingStation string // Code bureau de vote

	Position
}

type Tag struct {
	ID int64
}

type Contact struct {
	ID          int64
	Firstname   string `sql:"not null"`
	Surname     string `sql:"not null"`
	MarriedName string `db:"married_name"`
	Gender      string
	Birthdate   time.Time

	Mail   string
	Phone  string
	Mobile string

	Address *Address
	Tags    []Tag
}

func (c *Contact) NewRecord(db *sqlx.DB) error {
	result, err := db.NamedExec("INSERT INTO contacts (firstname, surname, married_name, gender, birthdate, mail, phone, mobile) VALUES (:firstname, :surname, :married_name, :gender, :birthdate, :mail, :phone, :mobile)", c)
	if err != nil {
		return err
	}
	c.ID, err = result.LastInsertId()
	return err
}

func (c *Contact) Update(db *sqlx.DB) error {
	return nil
}

func (c *Contact) Delete(db *sqlx.DB) error {
	return nil
}

func FindAllContacts(db *sqlx.DB) []Contact {
	contacts := []Contact{}
	if err := db.Select(&contacts, "SELECT id, firstname, lastname FROM contacts ORDER BY lastname DESC"); err != nil {
		fmt.Println(err)
	}
	return contacts
}

func FindContactByID(db *sqlx.DB, id int) *Contact {
	contact := Contact{}
	if err := db.Get(&contact, "SELECT * FROM contacts WHERE id=? LIMIT 1", id); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Println(err)
	}
	return &contact
}
