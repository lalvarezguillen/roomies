package room

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/helpers"
	"github.com/stretchr/testify/assert"
)

var testRoom = Room{
	ID:          "test-room",
	Title:       "New Test Room",
	Description: "Testing",
}
var jsonTestRoom, err = json.Marshal(testRoom)

func TestListEmptyRoomsColl(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/rooms/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleList(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody RoomsQueryResult
		json.Unmarshal(res.Body.Bytes(), respBody)
		assert.Empty(t, respBody.Rooms)
		assert.Equal(t, "", respBody.LastID)
	}
}

func TestPublishRoom(t *testing.T) {
	// setup
	defer helpers.ClearCollection(Collection)
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/rooms/",
		strings.NewReader(string(jsonTestRoom)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleCreate(c)) {
		assert.Equal(t, 201, res.Code)
		var respBody Room
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.Equal(t, testRoom.Title, respBody.Title)
		assert.Equal(t, testRoom.Description, respBody.Description)
	}
}

func TestListRooms(t *testing.T) {
	// setup
	NewRoom(&testRoom)
	defer helpers.ClearCollection(Collection)
	req := httptest.NewRequest(echo.GET, "/rooms/", strings.NewReader(""))
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, HandleList(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody RoomsQueryResult
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.NotEmpty(t, respBody.Rooms)
		assert.NotEmpty(t, respBody.LastID)
	}
}

func TestGetRoom(t *testing.T) {
	// setup
	newR, _ := NewRoom(&testRoom)
	defer helpers.ClearCollection(Collection)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/rooms/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(newR.ID)

	// test
	if assert.NoError(t, HandleGet(c)) {
		assert.Equal(t, 200, res.Code)
		var respRoom Room
		json.Unmarshal(res.Body.Bytes(), &respRoom)
		assert.Equal(t, newR.ID, respRoom.ID)
	}
}

func TestRemoveRoom(t *testing.T) {
	// setup
	newR, _ := NewRoom(&testRoom)
	defer helpers.ClearCollection(Collection)
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/rooms/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(newR.ID)

	// test
	if assert.NoError(t, HandleDelete(c)) {
		assert.Equal(t, 204, res.Code)
	}
}
