// Views for JSON responses
package views

import "github.com/quorumsco/territories/models"

// Territories is a type used for JSON request responses
type Territories struct {
	Territories []models.Territory `json:"territories"`
}

// Territory is a type used for JSON request responses
type Territory struct {
	Territory *models.Territory `json:"territory"`
}
