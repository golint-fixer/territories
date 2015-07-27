package models

import (
	"database/sql"

	"github.com/asaskevich/govalidator"
)

type Position struct {
	X float64
	Y float64

	Latitude  float64
	Longitude float64
}

type Address struct {
	ID uint

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

type Contact struct {
	ID          uint           `json:"id"`
	Firstname   string         `sql:"not null" json:"firstname"`
	Surname     string         `sql:"not null" json:"surname"`
	MarriedName sql.NullString `db:"married_name" json:"married_name,omitempty"`
	Gender      sql.NullString `json:"gender,omitempty"`
	Birthdate   NullTime       `json:"birthdate,omitempty"`
	Mail        sql.NullString `json:"mail,omitempty"`
	Phone       sql.NullString `json:"phone,omitempty"`
	Mobile      sql.NullString `json:"mobile,omitempty"`

	UserID uint

	Address *Address `json:"address,omitempty"`
	Notes   []Note   `json:"notes,omitempty"`
	Tags    []Tag    `json:"tags,omitempty"`
}

func (c *Contact) Validate() map[string]string {
	var errs = make(map[string]string)
	if c.Firstname == "" {
		errs["firstname"] = "is required"
	}
	if c.Surname == "" {
		errs["surname"] = "is required"
	}
	if c.Mail.Valid && !govalidator.IsEmail(c.Mail.String) {
		errs["mail"] = "is not valid"
	}
	return errs
}
