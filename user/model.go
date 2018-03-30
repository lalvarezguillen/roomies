package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

const Collection string = "users"

// User represents a user, either someone renting a Room, or
// someone looking for a Room, or someone who is or was a roommate of
// the Room
type User struct {
	gorm.Model
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"    gorm:"unique_index"`
	Phone     string    `json:"phone"`
	DOB       time.Time `json:"dob"`
	Bio       string    `json:"bio"`
	RoomID    int       `json:"roomID"`
}

// Users represents a list of User structs
type Users []User

// UsersListQuery Represents a query for User. Assits pagination
type UsersListQuery struct {
	LastID string
	Limit  int
}

// UsersQueryResult represents the result of a query for People.
type UsersQueryResult struct {
	Users  *Users
	LastID string
}
