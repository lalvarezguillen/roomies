package person

import "time"

const Collection string = "people"

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
	Bio              string
	RegistrationDate time.Time
}

// People represents a list of Person structs
type People []Person

// PeopleListQuery Represents a query for People. Assits pagination
type PeopleListQuery struct {
	LastID string
	Limit  int
}

// PeopleQueryResult represents the result of a query for People.
type PeopleQueryResult struct {
	People *People
	LastID string
}
