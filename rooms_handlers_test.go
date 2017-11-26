package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/config"
	"github.com/lalvarezguillen/roomies/room"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

var testRoom = room.Room{Title: "New Test Room",
	Description: "Testing"}
var jsonTestRoom, err = json.Marshal(testRoom)

func ClearCollection(collName string) {
	db := config.DB{}
	sess, err := db.DoDial()
	if err != nil {
		panic("There was a problem connecting to the DB")
	}
	defer sess.Close()
	coll := sess.DB(db.Name()).C(collName)
	coll.RemoveAll(bson.M{})
}

func InsertTestRoom(e *echo.Echo) {
	setupReq := httptest.NewRequest(echo.POST, "/rooms/",
		strings.NewReader(string(jsonTestRoom)))
	setupReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	setupRes := httptest.NewRecorder()
	setupCtx := e.NewContext(setupReq, setupRes)
	publishRoom(setupCtx)
}

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
	defer ClearCollection(room.Collection)
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/rooms/",
		strings.NewReader(string(jsonTestRoom)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// test
	if assert.NoError(t, publishRoom(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody room.Room
		json.Unmarshal(res.Body.Bytes(), &respBody)
		assert.Equal(t, testRoom.Title, respBody.Title)
		assert.Equal(t, testRoom.Description, respBody.Description)
	}
}

func TestListRooms(t *testing.T) {
	// setup
	//defer ClearCollection(room.Collection)
	e := echo.New()
	InsertTestRoom(e)
	defer ClearCollection(room.Collection)
	req := httptest.NewRequest(echo.GET, "/rooms/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	// test
	if assert.NoError(t, listRooms(c)) {
		assert.Equal(t, 200, res.Code)
		var respBody room.RoomsLastID
		json.Unmarshal(res.Body.Bytes(), &respBody)
		fmt.Println(respBody)
		assert.NotEmpty(t, respBody.Rooms)
		assert.NotEmpty(t, respBody.LastID)
	}
}
