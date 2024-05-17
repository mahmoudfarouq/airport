package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func itineraryHandler(c echo.Context) error {
	var request [][]string
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid request",
		})
	}

	if len(request) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "empty request",
		})
	}

	for _, tuple := range request {
		if len(tuple) != 2 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "invalid tuple",
				"tuple": tuple,
			})
		}
	}

	return c.JSON(http.StatusOK, computeItinerary(request))
}

func main() {
	e := echo.New()

	e.POST("/", itineraryHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
