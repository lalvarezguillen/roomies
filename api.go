package main

import (
	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/person"
	"github.com/lalvarezguillen/roomies/room"
)

func main() {
	api := echo.New()
	api.GET("/", getAPIInfo)

	r := api.Group("/rooms/")
	r.GET("", room.HandleList)
	r.GET(":id", room.HandleGet)
	r.POST("", room.HandleCreate)
	r.PUT(":id", room.HandleUpdate)
	r.DELETE(":id", room.HandleDelete)
	r.POST(":id/favorite", room.HandleFavorite)
	r.DELETE(":id/favorite", room.HandleUnfavorite)

	p := api.Group("/people/")
	p.GET("", person.HandleList)
	p.GET(":id", person.HandleGet)
	p.POST("", person.HandleCreate)
	p.PUT(":id", person.HandleUpdate)
	p.DELETE(":id", person.HandleDelete)

	api.Logger.Fatal(api.Start(":1234"))
}
