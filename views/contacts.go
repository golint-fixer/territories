// Views to communicate via JSON
package views

import "github.com/quorumsco/contacts/models"

type Contacts struct {
	Contacts []models.Contact `json:"contacts"`
}

type Contact struct {
	Contact *models.Contact `json:"contact"`
}
