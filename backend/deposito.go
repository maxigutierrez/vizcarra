package main

import (
	"net/http"

	"./config"
	NodosController "./controllers/nodos"
	UsuariosController "./controllers/usuarios"
	"./database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.Init()
	database.InitDb()

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

	a := e.Group("/api")
	b := e.Group("/api") //public

	a.Use(middleware.JWT([]byte("K*Xcu{0qGFViRGo7gGx7")))

	//Rutas
	//usuarios
	b.POST("/login", UsuariosController.LoginApp)
	a.GET("/authenticated", UsuariosController.GetAuthenticatedUser)

	a.POST("/lectura", NodosController.CrearLectura)
	a.GET("/lectura/:id", NodosController.GetLectura)
	a.GET("/destinos/:id", NodosController.GetDestinos)
	a.PUT("/lectura", NodosController.SetLectura)

	//e.Logger.Fatal(e.Start(config.PortDeposito))

	e.Logger.Fatal(e.StartTLS(config.PortDeposito, "/etc/letsencrypt/live/www.credifinexpress.com.ar/fullchain.pem", "/etc/letsencrypt/live/www.credifinexpress.com.ar/privkey.pem"))

}
