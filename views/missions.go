package views

import "github.com/quorumsco/contacts/models"

type Missions struct {
	Missions []models.Mission `json:"missions"`
}

type Mission struct {
	Mission *models.Mission `json:"mission"`
}
