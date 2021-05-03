package main

import (
	"net/http"

	TrackController "./controllers/track"
	Track_lecturasController "./controllers/track_lecturas"

	"./config"
	ClientesAutogestionController "./controllers/clientes_autogestion"
	EntregasController "./controllers/entregas"
	FranquiciasController "./controllers/franquicias"
	NodosController "./controllers/nodos"
	NoentregasController "./controllers/noentregas"
	ObservacionesController "./controllers/observaciones"
	RuteosController "./controllers/ruteos"
	UsuariosController "./controllers/usuarios"

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
	b := e.Group("/api") //public

	a.Use(middleware.JWT([]byte("K*Xcu{0qGFViRGo7gGx7")))

	//Rutas
	//usuarios
	b.POST("/login", UsuariosController.Login)
	b.POST("/clientes/autogestion/login", ClientesAutogestionController.Login)

	a.GET("/authenticated", UsuariosController.GetAuthenticatedUser)

	a.GET("/nodos/dominios", NodosController.Get_dominios)

	a.GET("/nodos/:id/auditoria", NodosController.Stock)
	a.POST("/nodos/auditoria", NodosController.Auditoria)

	//a.GET("/nodos/trasbordo/:dominio/:dominiodestino", NodosController.Trasbordo)
	//a.POST("/nodos/trasbordo", NodosController.Trasbordar)

	a.GET("/nodos/:id/cargar/:dominio", NodosController.Cargar)
	a.GET("/nodos/:id/descargar/:dominio", NodosController.Descargar)

	a.POST("/nodos/cargar", NodosController.CargarCamion)
	a.POST("/nodos/descargar", NodosController.DescargarCamion)

	a.GET("/nodos/:nodo/ingreso/:id", NodosController.Ingreso)
	a.POST("/nodos/ingreso", NodosController.IngresoFranquicia)

	a.GET("/nodos/:nodo/egreso/:id", NodosController.Egreso)
	a.POST("/nodos/egresos", NodosController.EgresoFranquicia)

	//ruteos
	a.GET("/recorrido", RuteosController.GetRecorridoCartero)
	a.POST("/recorrido", RuteosController.GetRecorrido)

	a.PUT("/ruteos/entrega", RuteosController.Entregar)
	a.GET("/noentregas", NoentregasController.GetAll)
	a.POST("/entregar", EntregasController.Entregar)

	a.GET("/franquicias", FranquiciasController.GetAll)

	b.GET("/track/:id", TrackController.ShowTrackCli)
	b.GET("/tracking/:id", TrackController.ShowTrackCli)
	a.GET("/track", Track_lecturasController.GetAllTracks)

	a.GET("/observaciones", ObservacionesController.GetAll)

	a.GET("/franquicias/atiende", FranquiciasController.GetFranquiciaZona)

	e.Logger.Fatal(e.Start(config.Port))

}
