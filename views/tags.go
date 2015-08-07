package views

import "github.com/quorumsco/contacts/models"

type Tags struct {
	Tags []models.Tag `json:"tags"`
}

type Tag struct {
	Tag *models.Tag `json:"tag"`
}
