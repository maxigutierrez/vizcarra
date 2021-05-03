package main

import (
	"net/http"

	ProductosController "./controllers/productos"

	"./config"
	"./database"
	l "./logs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.Init()
	database.InitDb()

	e := echo.New()
	e.Static("/static", "static")
	l.InfoLogger.Println("Iniciando..")
	l.AuditLogger.Println("Iniciando Audit..")
	l.SincLogger.Println("Iniciando..")
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

	// a := e.Group("/api")
	b := e.Group("/api") //public

	// b.Use(middleware.JWT([]byte("K*Xcu{0qGFViRGo7gGx7")))

	//Rutas

	//productos
	b.POST("/productos", ProductosController.Create)
	b.PUT("/productos", ProductosController.Update)
	b.GET("/productos", ProductosController.GetAll)
	b.GET("/productos/:id", ProductosController.Get)
	b.DELETE("/productos/:id", ProductosController.Delete)

	// e.Logger.Fatal(e.Start(config.Port))
	//e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(":5000"))

}
