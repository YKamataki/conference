package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Connect to DB
	db := ConnectDB()

	// Routes
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello World!\ngithub.com/YKamataki/conference")
	})

	e.GET("/api/conferences", func(ctx echo.Context) error {
		// Get all conferences
		conferences := GetConferences(db)
		// build json response
		resp, err := json.Marshal(conferences)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":443"))
}
