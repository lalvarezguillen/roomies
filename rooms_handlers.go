package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/room"
)

func listRooms(c echo.Context) error {
	roomsQ := room.RoomsListQuery{}
	err := c.Bind(roomsQ)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, "ERROR")
	}
	resp, err := room.ListRooms(&roomsQ)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return c.JSON(http.StatusOK, resp)
}

func getRoom(c echo.Context) error {
	return c.JSON(200, "")
}

func publishRoom(c echo.Context) error {
	return c.JSON(200, "")
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
