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
	resp := room.List(&roomsQ)
	fmt.Println(resp)
	return c.JSON(http.StatusOK, resp)
}

func getRoom(c echo.Context) error {
	r, err := room.GetByID(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(404, err)
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
	id := c.Param("id")
	var newData map[string]interface{}
	err := c.Bind(&newData)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, err)
	}
	updatedR, err := room.Update(id, &newData)
	if err != nil {
		c.JSON(404, err)
	}
	return c.JSON(200, updatedR)
}

func removeRoom(c echo.Context) error {
	err := room.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(404, "Room does not exists")
	}
	return c.JSON(200, "")
}

func favoriteRoom(c echo.Context) error {
	return c.JSON(200, "")
}

func unfavRoom(c echo.Context) error {
	return c.JSON(200, "")
}
