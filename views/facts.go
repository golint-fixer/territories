// Views for JSON responses
package views

import "github.com/quorumsco/contacts/models"

// Facts is a type used for JSON request responses
type Facts struct {
	Facts []models.Fact `json:"facts"`
}

// Fact is a type used for JSON request responses
type Fact struct {
	Fact *models.Fact `json:"fact"`
}

type FactsJson struct {
	FactsJson models.FactsJson `json:"facts_json"`
}
