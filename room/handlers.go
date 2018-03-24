package room

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleList(c echo.Context) error {
	roomsQ := RoomsListQuery{}
	err := c.Bind(&roomsQ)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	resp := List(&roomsQ)
	return c.JSON(http.StatusOK, resp)
}

func HandleGet(c echo.Context) error {
	r, err := GetByID(c.Param("id"))
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, r)
}

func HandleCreate(c echo.Context) error {
	r := Room{}
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(400, err)
	}
	New(&r)
	return c.JSON(201, r)
}

func HandleUpdate(c echo.Context) error {
	var updatedData Room
	err := c.Bind(&updatedData)
	if err != nil {
		return c.JSON(400, err)
	}
	updatedR, err := Update(&updatedData)
	if err != nil {
		c.JSON(404, err)
	}
	return c.JSON(200, updatedR)
}

func HandleDelete(c echo.Context) error {
	err := Delete(c.Param("id"))
	if err != nil {
		return c.JSON(404, "Room does not exists")
	}
	return c.JSON(204, nil)
}

func HandleFavorite(c echo.Context) error {
	return c.JSON(200, "")
}

func HandleUnfavorite(c echo.Context) error {
	return c.JSON(200, "")
}
