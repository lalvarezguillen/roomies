package person

import "time"

// Person represents a user, whether he's seeking or
// offering.
type Person struct {
	ID               string
	FirstName        string
	LastName         string
	Email            string
	Phone            string
	PwHash           string
	DOB              string
	Roommates        string
	Bio              string
	RegistrationDate time.Time
}

// People represents a list of Person structs
type People []Person
