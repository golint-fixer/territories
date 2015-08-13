package models

type Tag struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name" sql:"not null;unique"`
	Color uint   `json:"color"`
}
