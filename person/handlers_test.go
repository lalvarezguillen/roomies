package person

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lalvarezguillen/roomies/helpers"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var testPerson = Person{
	ID:        "test-person",
	FirstName: "Test",
	LastName:  "Person",
	Email:     "test@person.com",
	DOB:       "1990-01-01",
	Bio:       "A Person for testing",
}

var jsonTestPerson, err = json.Marshal(testPerson)

func TestListEmptyPeopleColl(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/people/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, HandleList(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody PeopleQueryResult
		json.Unmarshal(res.Body.Bytes(), respBody)
		assert.Empty(t, respBody.People)
		assert.Equal(t, "", respBody.LastID)
	}
}

func TestCreatePerson(t *testing.T) {
	// setup
	defer helpers.ClearCollection(Collection)
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/people/",
		strings.NewReader(string(jsonTestPerson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleCreate(c)) {
		assert.Equal(t, 201, res.Code)
		var respBody Person
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.Equal(t, testPerson.Email, respBody.Email)
		assert.Equal(t, testPerson.FirstName, respBody.FirstName)
	}
}

func TestListPeople(t *testing.T) {
	// setup
	defer helpers.ClearCollection(Collection)
	CreatePerson(&testPerson)
	req := httptest.NewRequest(echo.GET, "/people/", strings.NewReader(""))
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleList(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody PeopleQueryResult
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.NotEmpty(t, respBody.People)
		assert.NotEmpty(t, respBody.LastID)
	}
}

func TestGetPerson(t *testing.T) {
	// setup
	defer helpers.ClearCollection(Collection)
	newPerson, _ := CreatePerson(&testPerson)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/people/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(newPerson.ID)

	// test
	if assert.NoError(t, HandleGet(c)) {
		assert.Equal(t, 200, res.Code)
		var respPerson Person
		json.Unmarshal(res.Body.Bytes(), &respPerson)
		assert.Equal(t, newPerson.ID, respPerson.ID)
	}
}

func TestRemovePerson(t *testing.T) {
	// setup
	defer helpers.ClearCollection(Collection)
	newPerson, _ := CreatePerson(&testPerson)
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/people/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(newPerson.ID)

	// test
	if assert.NoError(t, HandleDelete(c)) {
		assert.Equal(t, 204, res.Code)
	}
}
