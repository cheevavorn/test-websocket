package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	A int  `json:"a"`
	B bool `json:"b"`
	C int  `json:"c"`
}

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write = Emit data to client
		err = ws.WriteMessage(websocket.TextMessage, []byte("Hello from server: "+time.Now().GoString()))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read = On Listening
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		data := Response{}
		err = json.Unmarshal([]byte(msg), &data)
		if err != nil {
			fmt.Printf("Cannot Unmarchal %s\n", msg)
		}
		fmt.Printf("msg: %v\n", data)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", ".")
	e.GET("/ws", hello) // ws://localhost:1323/ws
	e.Logger.Fatal(e.Start(":1323"))
}
