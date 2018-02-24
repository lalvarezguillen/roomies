package room

import (
	"github.com/lalvarezguillen/roomies/person"
)

const Collection string = "rooms"

// Rules represents the convivence rules set by the landlord.
type Rules struct {
	PetsOK     bool `json:"petsOK"`
	SmokingOK  bool `json:"smokingOK"`
	VisitorsOK bool `json:"visitorsOK"`
}

// Address holds the address of a room
type Address struct {
	Country    string `json:"country"`
	Region     string `json:"region"`
	City       string `json:"city"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Apartament string `json:"apartament"`
}

// Coordinates holds a latitude-longitude pair
type Coordinates struct {
	Lat  float32 `json:"lat"`
	Long float32 `json:"long"`
}

// Room represents a room on the market
type Room struct {
	ID               string          `json:"id"    bson:"_id,omitempty"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	RoommatesCount   int             `json:"roommatesCount"`
	Roommates        []person.Person `json:"roommates"`
	RegistrationDate int64           `json:"registrationDate"`
	Price            float32         `json:"price"`
	Rules            *Rules          `json:"rules"`
	Address          *Address        `json:"address"`
	Location         *Coordinates    `json:"location"`
	Media            []string        `json:"media"`
	Available        bool            `json:"available"`
}

// Rooms represents a list of rooms
type Rooms []Room

// RoomsListQuery represents the parameters of a query that requests a list of
// paginated rooms.
type RoomsListQuery struct {
	LastID string
	Limit  int
}

// RoomsQueryResult represents the results of a paginated query for Rooms
type RoomsQueryResult struct {
	Rooms  *Rooms
	LastID string
}
