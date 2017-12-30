package main

import (
	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/person"
)

func listPeople(c echo.Context) error {
	var peopleQ person.PeopleListQuery
	err := c.Bind(&peopleQ)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	resp := person.List(&peopleQ)
	return c.JSON(200, resp)
}

func getPerson(c echo.Context) error {
	p, err := person.GetByID(c.Param("id"))
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, p)
}

func createPerson(c echo.Context) error {
	var p person.Person
	err := c.Bind(&p)
	if err != nil {
		return c.JSON(400, err)
	}
	person.New(&p)
	return c.JSON(201, p)
}

func updatePerson(c echo.Context) error {
	return c.JSON(200, "")
}

func deletePerson(c echo.Context) error {
	return c.JSON(200, "")
}
