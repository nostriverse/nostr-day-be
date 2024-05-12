package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")
	
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app.Logger.Fatal(app.Start(port))
}