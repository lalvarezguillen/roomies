package user

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/config"
	"github.com/stretchr/testify/assert"
)

var dummyUser = User{
	FirstName: "Dummy",
	LastName:  "User",
	Email:     "dummy@user.com",
	DOB:       time.Now(),
	Bio:       "A User for testing",
}

var jsonTestUser, _ = json.Marshal(&dummyUser)

// func init() {
// 	config.DB.AutoMigrate(&User{})
// }

func TestListUsersEmpty(t *testing.T) {
	// setup
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, HandleList(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody Users
		json.Unmarshal(res.Body.Bytes(), respBody)
		assert.Empty(t, respBody)
	}
}

func TestHandleCreateUser(t *testing.T) {
	// setup
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/users/",
		strings.NewReader(string(jsonTestUser)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleCreate(c)) {
		assert.Equal(t, 201, res.Code)
		var respBody User
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.Equal(t, dummyUser.Email, respBody.Email)
		assert.Equal(t, dummyUser.FirstName, respBody.FirstName)
		var usersCount int
		config.DB.Model(&User{}).Count(&usersCount)
		assert.Equal(t, 1, usersCount)
	}
}

func TestHandleListUsers(t *testing.T) {
	// setup
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	config.DB.Create(&dummyUser)
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleList(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody Users
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.NotEmpty(t, respBody)
	}
}

func TestGetUser(t *testing.T) {
	// setup
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	config.DB.Create(&dummyUser)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(dummyUser.ID))

	// test
	if assert.NoError(t, HandleGet(c)) {
		assert.Equal(t, 200, res.Code)
		var responseUser User
		json.Unmarshal(res.Body.Bytes(), &responseUser)
		assert.Equal(t, dummyUser.ID, responseUser.ID)
	}
}

func TestRemovePerson(t *testing.T) {
	// setup
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	config.DB.Create(&dummyUser)
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(dummyUser.ID))

	// test
	if assert.NoError(t, HandleDelete(c)) {
		assert.Equal(t, 204, res.Code)
		var uc int
		config.DB.Model(&User{}).Count(&uc)
		assert.Equal(t, 0, uc)
	}
}

func TestUpdatePerson(t *testing.T) {
	// setup
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	config.DB.Create(&dummyUser)
	e := echo.New()
	var updatedData User
	copier.Copy(&updatedData, &dummyUser)
	updatedData.Email = "updated@email.com"
	jsonData, _ := json.Marshal(updatedData)
	req := httptest.NewRequest(echo.PUT, "/users/",
		strings.NewReader(string(jsonData)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(updatedData.ID))

	// test
	if assert.NoError(t, HandleUpdate(c)) {
		assert.Equal(t, 200, res.Code)
		var updatedUser User
		json.Unmarshal(res.Body.Bytes(), &updatedUser)
		assert.Equal(t, updatedData.Email, updatedUser.Email)
		var uc int
		config.DB.Model(&User{}).Count(&uc)
		assert.Equal(t, 1, uc)
	}
}

func TestUpdateNonexistentPerson(t *testing.T) {
	config.DB.AutoMigrate(&User{})
	defer config.DB.DropTable(&User{})
	e := echo.New()
	jsonData, _ := json.Marshal(dummyUser)
	req := httptest.NewRequest(echo.PUT, "/users/",
		strings.NewReader(string(jsonData)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// test
	if assert.NoError(t, HandleUpdate(c)) {
		assert.Equal(t, 404, res.Code)
	}
}
