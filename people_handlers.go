package main

import (
	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/person"
)

func listPeople(c echo.Context) error {
	peopleQ := person.PeopleListQuery{}
	err := c.Bind(&peopleQ)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	resp := person.List(&peopleQ)
	return c.JSON(200, resp)
}

func getPerson(c echo.Context) error {
	return c.JSON(200, "")
}

func createPerson(c echo.Context) error {
	return c.JSON(200, "")
}

func updatePerson(c echo.Context) error {
	return c.JSON(200, "")
}

func deletePerson(c echo.Context) error {
	return c.JSON(200, "")
}
