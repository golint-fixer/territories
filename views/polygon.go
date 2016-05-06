// Views for JSON responses
package views

import "github.com/quorumsco/territories/models"

// Notes is a type used for JSON request responses
type Polygon struct {
	Polygon []models.PointLocation `json:"polygon"`
}

// Note is a type used for JSON request responses
type PointLocation struct {
	PointLocation *models.PointLocation `json:"pointLocation"`
}
