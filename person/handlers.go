package person

import (
	"github.com/labstack/echo"
)

// HandleList deals with requests to query the People collection
func HandleList(c echo.Context) error {
	var peopleQ PeopleListQuery
	err := c.Bind(&peopleQ)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	resp := List(&peopleQ)
	return c.JSON(200, resp)
}

// HandleGet deals with requests to get a particular Person by ID
func HandleGet(c echo.Context) error {
	p, err := GetByID(c.Param("id"))
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, p)
}

// HandleCreate deals with requests to create a new Person entry
func HandleCreate(c echo.Context) error {
	var personData Person
	err := c.Bind(&personData)
	if err != nil {
		return c.JSON(400, err)
	}
	p, err := New(&personData)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(201, p)
}

// HandleUpdate deals with request to update a Person entry
func HandleUpdate(c echo.Context) error {
	var p Person
	err := c.Bind(&p)
	if err != nil {
		return c.JSON(400, err)
	}
	updatedP, err := Update(&p)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, &updatedP)
}

// HandleDelete deals with requests to delete a Person entry.
func HandleDelete(c echo.Context) error {
	p, err := GetByID(c.Param("id"))
	if err != nil {
		return c.JSON(404, "Does not exist")
	}
	Delete(p.ID)
	return c.JSON(204, nil)
}
