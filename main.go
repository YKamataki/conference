package main

import (
	"net/http"
  "json"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

  // Connect to DB 
		db := ConnectDB()
  
	e.GET("/api/conferences", func(ctx echo.Context) error {
    // Get all conferences
    conferences := GetConferences(db)
    // build json response
    resp, err := json.Marshal(conferences)
		return ctx.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":443"))
}


