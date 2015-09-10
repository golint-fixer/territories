package models

type Tag struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name" sql:"not null;unique"`
	Color string `json:"color"`
}

type TagArgs struct {
	GroupID   uint
	ContactID uint
	Tag       *Tag
}

type TagReply struct {
	Tag  *Tag
	Tags []Tag
}
