package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	A int  `json:"a"`
	B bool `json:"b"`
	C int  `json:"c"`
}

func main() {
	e := echo.New()

	hub := newHub()
	go hub.run()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", ".")
	e.GET("/ws", func(c echo.Context) error {
		ServeWs(hub, c)
		return nil
	})

	// e.GET("/ws", hello) // ws://localhost:1323/ws
	e.Logger.Fatal(e.Start(":1323"))
}
