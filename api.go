package main

import (
	"github.com/labstack/echo"
	"github.com/lalvarezguillen/roomies/user"
)

func main() {
	api := echo.New()
	// api.GET("/", getAPIInfo)

	// r := api.Group("/rooms/")
	// r.GET("", room.HandleList)
	// r.GET(":id", room.HandleGet)
	// r.POST("", room.HandleCreate)
	// r.PUT(":id", room.HandleUpdate)
	// r.DELETE(":id", room.HandleDelete)
	// r.POST(":id/favorite", room.HandleFavorite)
	// r.DELETE(":id/favorite", room.HandleUnfavorite)

	p := api.Group("/users/")
	p.GET("", user.HandleList)
	p.GET(":id", user.HandleGet)
	p.POST("", user.HandleCreate)
	p.PUT(":id", user.HandleUpdate)
	p.DELETE(":id", user.HandleDelete)

	api.Logger.Fatal(api.Start(":1234"))
}
