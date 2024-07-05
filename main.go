package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		db := ConnectDB()
		MigrateDB(db)
		return c.String(http.StatusOK, "DB connection and migration are successful!")
	})

	e.Logger.Fatal(e.Start(":443"))
}
