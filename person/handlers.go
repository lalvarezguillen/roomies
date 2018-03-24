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
	resp := ListPeople(&peopleQ)
	return c.JSON(200, resp)
}

// HandleGet deals with requests to get a particular Person by ID
func HandleGet(c echo.Context) error {
	p, err := GetPersonByID(c.Param("id"))
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
	p, err := CreatePerson(&personData)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(201, p)
}

// HandleUpdate deals with request to update a Person entry
func HandleUpdate(c echo.Context) error {
	var newData Person
	if err := c.Bind(&newData); err != nil {
		return c.JSON(400, err)
	}
	personID := c.Param("id")
	if newData.ID != personID {
		resp := map[string]string{"error": "The ID of a Person can't be changed"}
		return c.JSON(400, resp)
	}
	if _, err := GetPersonByID(personID); err != nil {
		resp := map[string]string{"error": "Does not exist"}
		return c.JSON(404, resp)
	}
	updatedP, err := UpdatePerson(&newData)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, &updatedP)
}

// HandleDelete deals with requests to delete a Person entry.
func HandleDelete(c echo.Context) error {
	personID := c.Param("id")
	if _, err := GetPersonByID(personID); err != nil {
		return c.JSON(404, "Does not exist")
	}
	DeletePerson(personID)
	return c.JSON(204, nil)
}
