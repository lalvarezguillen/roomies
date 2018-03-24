package person

import "time"

const Collection string = "people"

// Person represents a user, whether he's seeking or
// offering.
type Person struct {
	ID               string    `json:"id"    bson:"_id,omitempty"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	DOB              string    `json:"dob"`
	Bio              string    `json:"bio"`
	RegistrationDate time.Time `json:"registrationDate"`
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
