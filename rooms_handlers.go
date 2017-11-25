package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/room"
)

func listRooms(c echo.Context) error {
	roomsQ := room.RoomsListQuery{}
	err := c.Bind(&roomsQ)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, "Error parsing the request")
	}
	resp := room.ListRooms(&roomsQ)
	fmt.Println(resp)
	return c.JSON(http.StatusOK, resp)
}

func getRoom(c echo.Context) error {
	r, err := room.GetByID(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(404, "The room requested does not exist")
	}
	return c.JSON(200, r)
}

func publishRoom(c echo.Context) error {
	r := room.Room{}
	err := c.Bind(&r)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, err)
	}
	room.New(&r)
	return c.JSON(200, r)
}

func updateRoom(c echo.Context) error {
	return c.JSON(200, "")
}

func removeRoom(c echo.Context) error {
	return c.JSON(200, "")
}

func favoriteRoom(c echo.Context) error {
	return c.JSON(200, "")
}

func unfavRoom(c echo.Context) error {
	return c.JSON(200, "")
}
