package views

import "github.com/Quorumsco/contact/models"

type ContactsView struct {
	Contacts []models.Contact `json:"contacts"`
}

type ContactView struct {
	Contact *models.Contact `json:"contact"`
}
