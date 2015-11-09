// Views for JSON responses
package views

import "github.com/quorumsco/contacts/models"

// Tags is a type used for JSON request responses
type Tags struct {
	Tags []models.Tag `json:"tags"`
}

// Tag is a type used for JSON request responses
type Tag struct {
	Tag *models.Tag `json:"tag"`
}
