// Definition of the structures and SQL interaction functions
package models

// Address represents all the components of a contact's address
type PointLocation struct {
	ID          uint   `db:"id" json:"id"`
	GroupID     uint   `db:"group_id" json:"group_id"`
	TerritoryID uint   `db:"territory_id" json:"territory_id"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	Location    string `json:"location,omitempty"` // as "lat,lon" (for elasticsearch)
}

// ContactArgs is used in the RPC communications between the gateway and Contacts
type PointLocationArgs struct {
	GroupID       uint
	TerritoryID   uint
	PointLocation *PointLocation
}

// ContactReply is used in the RPC communications between the gateway and Contacts
type PointLocationReply struct {
	PointLocation *PointLocation
	Polygon       []PointLocation
}
