// Views for JSON responses
package views

import "github.com/quorumsco/contacts/models"

// Missions is a type used for JSON request responses
type Missions struct {
	Missions []models.Mission `json:"missions"`
}

// Mission is a type used for JSON request responses
type Mission struct {
	Mission *models.Mission `json:"mission"`
}
