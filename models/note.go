package models

import "time"

type Note struct {
	ID      uint
	Content string
	Date    time.Time

	UserID    uint `db:"user_id" json:"user_id"`
	ContactID uint `db:"contact_id" json:"contact_id"`
}
