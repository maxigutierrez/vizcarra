package main

import (
	"net/http"

	RolesController "./controllers_liquidaciones/roles"
	Clientes_ComisionesController "./controllers_liquidaciones/clientes_comisiones"
	PartesController "./controllers_liquidaciones/partes"
	Partes_franquiciasController "./controllers_liquidaciones/partes_franquicias"

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

	a := e.Group("/api")
	// b := e.Group("/api") //public

	a.Use(middleware.JWT([]byte("K*Xcu{0qGFViRGo7gGx7")))

	//roles
	a.POST("/roles", RolesController.Create)
	a.PUT("/roles", RolesController.Update)
	a.GET("/roles", RolesController.GetAll)
	a.GET("/roles/:id", RolesController.Get)
	a.DELETE("/roles/:id", RolesController.Delete)

	//clientes_comisiones
	a.POST("/clientes_comisiones", Clientes_ComisionesController.Create)
	a.PUT("/clientes_comisiones", Clientes_ComisionesController.Update)
	a.GET("/clientes_comisiones", Clientes_ComisionesController.GetAll)
	a.GET("/clientes_comisiones/:id", Clientes_ComisionesController.Get)
	a.DELETE("/clientes_comisiones/:id", Clientes_ComisionesController.Delete)

	//partes
	a.POST("/partes", PartesController.Create)
	a.PUT("/partes", PartesController.Update)
	a.GET("/partes", PartesController.GetAll)
	a.GET("/partes/view", PartesController.GetAllView)
	a.GET("/partes/:id", PartesController.Get)
	a.DELETE("/partes/:id", PartesController.Delete)

	//partes_franquicias
	a.POST("/partes_franquicias", Partes_franquiciasController.Create)
	a.PUT("/partes_franquicias", Partes_franquiciasController.Update)
	a.GET("/partes_franquicias", Partes_franquiciasController.GetAll)
	a.GET("/partes_franquicias/view", Partes_franquiciasController.GetAllView)
	a.GET("/partes_franquicias/:id", Partes_franquiciasController.Get)
	a.DELETE("/partes_franquicias/:id", Partes_franquiciasController.Delete)

	e.Logger.Fatal(e.Start(config.PortLiq))
	// e.Logger.Fatal(e.Start(":8000"))
}
