package main

import "github.com/labstack/echo"

func main() {
	api := echo.New()
	api.GET("/", getAPIInfo)

	api.GET("/rooms", listRooms)
	api.GET("/rooms/:id", getRoom)
	api.POST("/rooms/:id", publishRoom)
	api.PUT("/rooms/:id", updateRoom)
	api.DELETE("/rooms/:id", removeRoom)
	api.POST("/rooms/:id/favorite", favoriteRoom)
	api.DELETE("/rooms/:id/favorite", unfavRoom)

	api.GET("/people", listPeople)
	api.GET("/people/:id", getPerson)
	api.POST("/people/:id", createPerson)
	api.PUT("/people/:id", updatePerson)
	api.DELETE("/people/:id", deletePerson)

	api.Logger.Fatal(api.Start(":1234"))
}
