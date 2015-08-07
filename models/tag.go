package models

type Tag struct {
	ID    uint
	Name  string `json:"name" sql:"not null;unique"`
	Color uint   `json:"color"`
}
