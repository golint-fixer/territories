// Views for JSON responses
package views

import "github.com/quorumsco/contacts/models"

// Notes is a type used for JSON request responses
type Notes struct {
	Notes []models.Note `json:"notes"`
}

// Note is a type used for JSON request responses
type Note struct {
	Note *models.Note `json:"note"`
}
