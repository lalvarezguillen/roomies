package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/room"
	"github.com/stretchr/testify/assert"
)

func TestListEmptyRoomsColl(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/rooms/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, listRooms(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody room.RoomsLastID
		json.Unmarshal(res.Body.Bytes(), respBody)
		assert.Empty(t, respBody.Rooms)
		assert.Equal(t, "", respBody.LastID)
	}
}

func TestPublishRoom(t *testing.T) {
	// setup
	e := echo.New()
	newR := room.Room{Title: "New Test Room", Description: "Testing"}
	jsonRoom, err := json.Marshal(newR)
	if err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(echo.POST, "/rooms/",
		strings.NewReader(string(jsonRoom)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, publishRoom(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody room.Room
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.Equal(t, newR.Title, respBody.Title)
		assert.Equal(t, newR.Description, respBody.Description)
	}
}
