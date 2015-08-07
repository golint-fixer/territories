package models

type Tag struct {
	ID    uint
	Name  string `json:"name"`
	Color uint   `json:"color"`
}
