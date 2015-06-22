package views

import "../models"

type Contacts struct {
	Contacts []models.Contact `json:"contacts"`
}

type Contact struct {
	Contact *models.Contact `json:"contact"`
}
