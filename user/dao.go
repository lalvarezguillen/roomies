package user

import (
	"github.com/lalvarezguillen/roomies/config"
)

// ListUsers fetches paginated people
func ListUsers(q *UsersQuery) UsersQueryResult {
	var us Users
	config.DB.Where(&q.Filters).Offset(q.Offset).Limit(q.Limit).Find(&us)
	return UsersQueryResult{&us}
}

// GetUserByID obtains a Person by ID
func GetUserByID(id string) (*User, error) {
	var u User
	config.DB.Where("id = ?", id).First(&u)
	return &u, nil
}

// NewUser creates a new User
func NewUser(u *User) (*User, error) {
	config.DB.Create(u)
	return u, nil
}

// DeleteUser removes a User
func DeleteUser(id string) error {
	var u User
	config.DB.Where("id = ?", id).First(&u)
	config.DB.Delete(&u)
	return nil
}

// UpdateUser updates a User's DB entry
func UpdateUser(u *User) (*User, error) {
	config.DB.Save(u)
	return u, nil
}
