package room

import (
	"github.com/lalvarezguillen/roomies/person"
)

const collection string = "rooms"

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

// MediaFile represents images and videos that can be associated to rooms
// or users
type MediaFile struct {
	ID           string `json:"id"`
	MimeType     string `json:"mimeType"`
	ThumbnailURL string `json:"thumbnailURL"`
	URL          string `json:"url"`
}

// Room represents a room on the market
type Room struct {
	ID               string          `json:"id"    bson:"_id,omitempty"`
	Ttle             string          `json:"title"`
	Description      string          `json:"description"`
	RoommatesCount   int             `json:"roommatesCount"`
	Roommates        []person.Person `json:"roommates"`
	RegistrationDate int64           `json:"registrtionDate"`
	Price            float32         `json:"price"`
	Rules            Rules           `json:"rules"`
	Address          Address         `json:"address"`
	Location         Coordinates     `json:"location"`
	Media            []MediaFile     `json:"media"`
	Available        bool            `json:"available"`
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
