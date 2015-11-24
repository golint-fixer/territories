// Definition of the structures and SQL interaction functions
package models

// Address represents all the components of a contact's address
type Action struct {
	ID uint `json:"id"`

	ActionID uint   `json:"action_id"`
	TypeData string `json:"type_data"`
	Data     string `json:"data"`
	Pitch    string `json:"pitch"`
	Status   string `json:"status"`
}

// Contact represents all the components of a contact
type Fact struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	GroupID  uint    `json:"group_id"`
	Date     string  `json:"-"`
	Type     string  `json:"type"`
	Status   string  `json:"status"`
	Contact  Contact `json:"contact"`
	Action   Action  `json:"action"`
	ActionID uint    `json:"-"`
}

// FactArgs is used in the RPC communications between the gateway and Contacts
type FactArgs struct {
	Fact *Fact
}

// FactReply is used in the RPC communications between the gateway and Contacts
type FactReply struct {
	Fact  *Fact
	Facts []Fact
}
