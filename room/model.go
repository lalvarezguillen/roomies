package room

import (
	"time"

	"github.com/lalvarezguillen/roomies/person"
)

const coll string = "rooms"

// Rules represents the convivence rules set by the landlord.
type Rules struct {
	PetsOK     bool
	SmokingOK  bool
	VisitorsOK bool
}

// Address holds the address of a room
type Address struct {
	Country    string
	Region     string
	City       string
	Street     string
	Number     string
	Apartament string
}

// Coordinates holds a latitude-longitude pair
type Coordinates struct {
	Lat  float32
	Long float32
}

// MediaFile represents images and videos that can be associated to rooms
// or users
type MediaFile struct {
	ID           string
	MimeType     string
	ThumbnailURL string
	URL          string
}

// Room represents a room on the market
type Room struct {
	ID               string
	Ttle             string
	Description      string
	RoommatesCount   int
	Roomates         []person.Person
	RegistrationDate time.Time
	Price            float32
	Rules            Rules
	Address          Address
	Location         Coordinates
	Media            []MediaFile
	Available        bool
}

// Rooms represents a list of rooms
type Rooms []Room

type RoomsListQuery struct {
	LastID string
	Limit  int
}

type RoomsLastID struct {
	Rooms  *Rooms
	LastID string
}
