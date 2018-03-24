package room

import (
	"net/http"

	"github.com/labstack/echo"
)

// HandleList deals with requests to query Rooms
func HandleList(c echo.Context) error {
	roomsQ := RoomsListQuery{}
	err := c.Bind(&roomsQ)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	resp := ListRooms(&roomsQ)
	return c.JSON(http.StatusOK, resp)
}

// HandleGet deals with requests to get a Room by its ID
func HandleGet(c echo.Context) error {
	r, err := GetRoomByID(c.Param("id"))
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, r)
}

// HandleCreate deals with requests to create a new Room
func HandleCreate(c echo.Context) error {
	r := Room{}
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(400, err)
	}
	CreateRoom(&r)
	return c.JSON(201, r)
}

// HandleUpdate deals with requests to update a particular Room
func HandleUpdate(c echo.Context) error {
	var r Room
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(400, err)
	}
	if r.ID != c.Param("id") {
		resp := map[string]string{"error": "The ID of a Room can't be changed"}
		return c.JSON(400, resp)
	}
	updatedR, err := UpdateRoom(&r)
	if err != nil {
		c.JSON(404, err)
	}
	return c.JSON(200, updatedR)
}

// HandleDelete deals with requests to update a particular Room's entry
func HandleDelete(c echo.Context) error {
	err := DeleteRoom(c.Param("id"))
	if err != nil {
		return c.JSON(404, "Room does not exists")
	}
	return c.JSON(204, nil)
}

// HandleFavorite deals with request to favorite a Room by a Person
func HandleFavorite(c echo.Context) error {
	return c.JSON(200, "")
}

// HandleUnfavorite deals with request to remove a favorite relationship between
// a Person and Room
func HandleUnfavorite(c echo.Context) error {
	return c.JSON(200, "")
}
