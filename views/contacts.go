package views

import "github.com/Quorumsco/contacts/models"

type Contacts struct {
	Contacts []models.Contact `json:"contacts"`
}

type Contact struct {
	Contact *models.Contact `json:"contact"`
}
