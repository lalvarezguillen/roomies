package user

// func init() {
// 	config.DB.AutoMigrate(&User{})
// }

// func TestCreateUser(t *testing.T) {
// 	defer config.DB.DropTable(&User{})
// 	testUser := User{
// 		FirstName: "Test",
// 		LastName:  "User",
// 		Email:     "test@user.com",
// 		Phone:     "0800user",
// 		DOB:       time.Now(),
// 		Bio:       "This is a test user",
// 	}
// 	assert.True(t, config.DB.NewRecord(testUser))
// 	config.DB.Create(&testUser)
// 	assert.False(t, config.DB.NewRecord(testUser))

// 	var usersCount int
// 	var existingUsers Users
// 	config.DB.Find(&existingUsers).Count(&usersCount)
// 	assert.Equal(t, usersCount, 1)
// }
