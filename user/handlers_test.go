package user

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

var jsonTestUser, _ = json.Marshal(dummyUser)

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
	}
}

func TestListPeople(t *testing.T) {
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

// func TestGetPerson(t *testing.T) {
// 	// setup
// 	defer helpers.ClearCollection(Collection)
// 	newPerson, _ := NewPerson(&testPerson)
// 	e := echo.New()
// 	req := httptest.NewRequest(echo.GET, "/people/", strings.NewReader(""))
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("id")
// 	c.SetParamValues(newPerson.ID)

// 	// test
// 	if assert.NoError(t, HandleGet(c)) {
// 		assert.Equal(t, 200, res.Code)
// 		var respPerson Person
// 		json.Unmarshal(res.Body.Bytes(), &respPerson)
// 		assert.Equal(t, newPerson.ID, respPerson.ID)
// 	}
// }

// func TestRemovePerson(t *testing.T) {
// 	// setup
// 	defer helpers.ClearCollection(Collection)
// 	newPerson, _ := NewPerson(&testPerson)
// 	e := echo.New()
// 	req := httptest.NewRequest(echo.DELETE, "/people/", strings.NewReader(""))
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("id")
// 	c.SetParamValues(newPerson.ID)

// 	// test
// 	if assert.NoError(t, HandleDelete(c)) {
// 		assert.Equal(t, 204, res.Code)
// 	}
// }

// func TestUpdatePerson(t *testing.T) {
// 	// setup
// 	defer helpers.ClearCollection(Collection)
// 	newPerson, _ := NewPerson(&testPerson)
// 	e := echo.New()
// 	updatedData := newPerson
// 	updatedData.Email = "updated@email.com"
// 	jsonData, _ := json.Marshal(updatedData)
// 	req := httptest.NewRequest(echo.PUT, "/people/",
// 		strings.NewReader(string(jsonData)))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("id")
// 	c.SetParamValues(newPerson.ID)

// 	// test
// 	if assert.NoError(t, HandleUpdate(c)) {
// 		assert.Equal(t, 200, res.Code)
// 		var updatedPerson Person
// 		json.Unmarshal(res.Body.Bytes(), &updatedPerson)
// 		assert.Equal(t, updatedPerson.Email, updatedData.Email)
// 	}
// }

// func TestUpdatePersonOverwritingID(t *testing.T) {
// 	defer helpers.ClearCollection(Collection)
// 	newPerson, _ := NewPerson(&testPerson)
// 	e := echo.New()
// 	var updatedData Person
// 	copier.Copy(&updatedData, &newPerson)
// 	updatedData.Email = "updated@email.com"
// 	updatedData.ID = "overwriting-id"
// 	jsonData, _ := json.Marshal(updatedData)
// 	req := httptest.NewRequest(echo.PUT, "/people/",
// 		strings.NewReader(string(jsonData)))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("id")
// 	c.SetParamValues(newPerson.ID)

// 	fmt.Println(updatedData.ID)
// 	fmt.Println(newPerson.ID)
// 	// test
// 	if assert.NoError(t, HandleUpdate(c)) {
// 		assert.Equal(t, 400, res.Code)
// 		fmt.Println(string(res.Body.Bytes()))

// 	}
// }

// func TestUpdateNonexistentPerson(t *testing.T) {
// 	defer helpers.ClearCollection(Collection)
// 	newPerson, _ := NewPerson(&testPerson)
// 	e := echo.New()
// 	updatedData := newPerson
// 	updatedData.Email = "updated@email.com"
// 	updatedData.ID = "nonexistent-person"
// 	jsonData, _ := json.Marshal(updatedData)
// 	req := httptest.NewRequest(echo.PUT, "/people/",
// 		strings.NewReader(string(jsonData)))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("id")
// 	c.SetParamValues("nonexistent-person")

// 	// test
// 	if assert.NoError(t, HandleUpdate(c)) {
// 		assert.Equal(t, 404, res.Code)
// 	}
// }
