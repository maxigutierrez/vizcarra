package main

import (
	"net/http"
	UsuariosController "./controllers/usuarios"
	SincroController "./controllers/sincro"
  ComprobantesController "./controllers/comprobantes"


	l "./logs"
	"./database"
  "./config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.Init()
	database.InitDb()

	e := echo.New()
	e.Static("/static", "static")
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

	a := e.Group("/api")
	b := e.Group("/api") //public

	a.Use(middleware.JWT([]byte("K*Xcu{0qGFViRGo7gGx7")))

	//Rutas
	//usuarios
	b.POST("/login", UsuariosController.Login)

	//Sincronizaciones Desktop
	a.POST("/sincro/config", SincroController.GetConfigurar)
	a.POST("/sincro/usuarios", SincroController.GetUsuarios)
	a.POST("/sincro/franquicias", SincroController.GetFranquicias)
	a.POST("/sincro/articulos", SincroController.GetArticulos)
	a.POST("/sincro/clientes", SincroController.GetClientes)
	a.POST("/sincro/localidades", SincroController.GetLocalidades)
	a.POST("/sincro/scripts", SincroController.GetScripts)
	a.POST("/sincro/comprobantes", ComprobantesController.Create)



	// Sincro APP
	a.GET("/sincro/lineas", SincroController.GetLineas_sincro)
	a.GET("/sincro/obs", SincroController.GetObs)
	a.GET("/sincro/nodos", SincroController.GetNodos_sincro)
	a.POST("/sincro/clientes/new", SincroController.SincroNew)
	a.POST("/sincro/redespachos", SincroController.SincroRedespachos)


	e.Logger.Fatal(e.Start(config.PortSincro))

}
