package main

import (
	"net/http"

	AfipController "./controllers/afip"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	b := e.Group("/api") //public

	b.GET("/afip/cuit/:id", AfipController.GetCUIT)
	e.GET("/afip/cuit/:id", AfipController.GetCUIT)
	e.Logger.Fatal(e.Start(":5600"))
	//e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
