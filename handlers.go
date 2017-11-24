package main

import "github.com/labstack/echo"

func getAPIInfo(c echo.Context) error {
	return c.JSON(200, "")
}
