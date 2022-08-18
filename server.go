package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	hub := newHub()
	go hub.run()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", ".")
	e.GET("/ws", func(c echo.Context) error {
		ServeWs(hub, c) // ws://localhost:1323/ws
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
