// Definition of the structures and SQL interaction functions
package models

// Contact represents all the components of a contact
type Territory struct {
	ID      uint            `gorm:"primary_key" json:"id"`
	Name    string          `sql:"not null" json:"name"`
	GroupID uint            `sql:"not null" db:"group_id" json:"-"`
	Polygon []PointLocation `json:"polygon,omitempty"`
}

//To represent a GeoPolygon
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// ContactArgs is used in the RPC communications between the gateway and Contacts
type TerritoryArgs struct {
	GroupID   uint
	Territory *Territory
}

// ContactReply is used in the RPC communications between the gateway and Contacts
type TerritoryReply struct {
	Territory   *Territory
	Territories []Territory
}

// Validate checks if the contact is valid
func (c *Territory) Validate() map[string]string {
	var errs = make(map[string]string)

	// if c.Firstname == "" {
	// 	errs["firstname"] = "is required"
	// }

	// if c.Surname == "" {
	// 	errs["surname"] = "is required"
	// }

	// if c.Mail != nil && !govalidator.IsEmail(*c.Mail) {
	// 	errs["mail"] = "is not valid"
	// }

	return errs
}
