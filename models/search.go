package models

type Search struct {
	Query string
	Field string
}

type SearchArgs struct {
	Search *Search
}

type SearchReply struct {
	Contacts []Contact
}
