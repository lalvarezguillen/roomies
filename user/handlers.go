package user

import (
	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/config"
)

// HandleList deals with requests to query the People collection
func HandleList(c echo.Context) error {
	var q UsersQuery
	err := c.Bind(&q)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	var us Users
	config.DB.Where(&q.Filters).Offset(q.Offset).Limit(q.Limit).Find(&us)
	return c.JSON(200, &us)
}

// HandleGet deals with requests to get a particular Person by ID
func HandleGet(c echo.Context) error {
	id := c.Param("id")
	var u User
	config.DB.Where("id = ?", id).First(&u)
	return c.JSON(200, &u)
}

// HandleCreate deals with requests to create a new Person entry
func HandleCreate(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(400, err)
	}
	config.DB.Create(u)
	return c.JSON(201, &u)
}

// HandleUpdate deals with request to update a Person entry
func HandleUpdate(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(400, "Error parsing the request")
	}
	config.DB.Save(u)
	return c.JSON(200, &u)
}

// HandleDelete deals with requests to delete a Person entry.
func HandleDelete(c echo.Context) error {
	id := c.Param("id")
	var u User
	config.DB.Where("id = ?", id).First(&u)
	config.DB.Delete(&u)
	return c.JSON(204, nil)
}
