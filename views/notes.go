package views

import "github.com/quorumsco/contacts/models"

type Notes struct {
	Notes []models.Note `json:"notes"`
}

type Note struct {
	Note *models.Note `json:"note"`
}
