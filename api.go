package main

import "github.com/labstack/echo"

func main() {
	api := echo.New()
	api.GET("/", getAPIInfo)

	r := api.Group("/rooms")
	r.GET("", listRooms)
	r.GET(":id", getRoom)
	r.POST(":id", publishRoom)
	r.PUT(":id", updateRoom)
	r.DELETE(":id", removeRoom)
	r.POST(":id/favorite", favoriteRoom)
	r.DELETE(":id/favorite", unfavRoom)

	p := api.Group("/people")
	p.GET("", listPeople)
	p.GET(":id", getPerson)
	p.POST("", createPerson) // We may not need this
	p.PUT(":id", updatePerson)
	p.DELETE(":id", deletePerson)

	api.Logger.Fatal(api.Start(":1234"))
}
