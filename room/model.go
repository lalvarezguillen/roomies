package room

import (
	"time"

	"github.com/lalvarezguillen/roomies/person"
)

// Room represents a room on the market
type Room struct {
	ID               string
	Ttle             string
	Description      string
	RoommatesCount   int
	Roomates         []person.Person
	PetsOK           bool
	SmokingOK        bool
	RegistrationDate time.Time
	Price            float32
}

// Rooms represents a list of rooms
type Rooms []Room
