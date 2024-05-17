package main

import (
	"github.com/labstack/echo/v4"
)

func itineraryHandler(c echo.Context) error {
	var request [][]string
	if err := c.Bind(&request); err != nil {
		return err
	}

	return c.JSON(200, request)
}

func main() {
	e := echo.New()

	e.POST("/", itineraryHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
