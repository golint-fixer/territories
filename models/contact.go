package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Position struct {
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
	ID          int64      `json:"id"`
	Firstname   *string    `sql:"not null" json:"firstname"`
	Surname     *string    `sql:"not null" json:"surname"`
	MarriedName *string    `db:"married_name" json:"married_name,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Birthdate   *time.Time `json:"birthdate,omitempty"`

	Mail   *string `json:"mail,omitempty"`
	Phone  *string `json:"phone,omitempty"`
	Mobile *string `json:"mobile,omitempty"`

	Address *Address `json:"address,omitempty"`
	Tags    *[]Tag   `json:"tags,omitempty"`
}

func (c *Contact) Validate() map[string]string {
	var errs = make(map[string]string)
	if c.Firstname == nil {
		errs["firstname"] = "is required"
	}
	if c.Surname == nil {
		errs["surname"] = "is required"
	}
	if c.Mail != nil && !govalidator.IsEmail(*c.Mail) {
		errs["mail"] = "is not valid"
	}
	return errs
}
