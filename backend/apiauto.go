package main

import (
	"net/http"

	AfipController "./controllers/afip"
	AgendaController "./controllers/agenda"
	ArticulosController "./controllers/articulos"
	AutogestionController "./controllers/autogestion"
	ClientesController "./controllers/clientes"
	ClientesAutogestionController "./controllers/clientes_autogestion"
	ClientesContactosController "./controllers/clientes_contactos"
	ClientesConveniosController "./controllers/clientes_convenios"
	ClientesDestinosController "./controllers/clientes_destinos"
	ClientesDomiciliosController "./controllers/clientes_domicilios"
	ComprobantesController "./controllers/comprobantes"
	ComprobantesTiposController "./controllers/comprobantes_tipos"
	CondicionivaController "./controllers/condicioniva"
	DocumentotiposController "./controllers/documentotipos"
	Domicilios_tiposController "./controllers/domicilios_tipos"
	EmpresasController "./controllers/empresas"
	EntregasController "./controllers/entregas"
	FranquiciasController "./controllers/franquicias"
	FranquiciasLocalidadesController "./controllers/franquicias_localidades"
	FrecuenciasController "./controllers/frecuencias"
	LocalidadesController "./controllers/localidades_ar"
	ProvinciasController "./controllers/provincias"
	RangosHorariosController "./controllers/rangos_horarios"
	RecibosController "./controllers/recibos"
	RetirosController "./controllers/retiros"
	RolesController "./controllers/roles"
	ServiciosController "./controllers/servicios"
	SistemasController "./controllers/sistemas"
	TrackController "./controllers/track"
	Track_lecturasController "./controllers/track_lecturas"
	UsuariosController "./controllers/usuarios"

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
	b := e.Group("/api") //public

	a.Use(middleware.JWT([]byte("K*Xcu{0qGFViRGo7gGx7")))

	//Rutas
	//usuarios
	b.POST("/login", UsuariosController.Login)
	a.PUT("/password/:id", UsuariosController.Password)
	a.PUT("/blanqueo/:id", UsuariosController.Blanqueo)
	a.GET("/authenticated", UsuariosController.GetAuthenticatedUser)
	a.PUT("/usuarios", UsuariosController.Update)
	a.PUT("/usuarios/baja/:id", UsuariosController.Baja)
	a.PUT("/usuarios/alta/:id", UsuariosController.Alta)
	a.GET("/usuarios/:id", UsuariosController.Get)
	a.POST("/usuarios", UsuariosController.Create)
	a.GET("/usuarios", UsuariosController.GetAll)
	a.GET("/usuarios/franquicias", FranquiciasController.GetFranquiciasbyusuario)
	a.GET("/usuarios/roles", RolesController.GetRolesbyusuario)
	a.GET("/usuarios/sistemas", SistemasController.GetSistemasbyusuario)

	a.GET("/roles", RolesController.GetAll)

	//franquicias
	a.GET("/franquicias/all", FranquiciasController.GetAllconNodos)
	a.GET("/franquicias/:id", FranquiciasController.Get)
	a.PUT("/franquicias", FranquiciasController.Update)
	a.POST("/franquicias", FranquiciasController.Create)
	a.GET("/franquicias", FranquiciasController.GetAll)
	a.DELETE("/franquicias/:id", FranquiciasController.Delete)
	a.GET("/franquicias/atiende", FranquiciasController.GetFranquiciaZona)
	a.GET("/franquicias/noatiende", FranquiciasController.GetFranquiciaZonaExcluir)
	a.GET("/franquicias/red", FranquiciasController.GetAllRed)

	a.POST("/recibos", RecibosController.Create)

	a.POST("/cobertura", FranquiciasLocalidadesController.GetCobertura)

	b.POST("/clientes/autogestion/login", ClientesAutogestionController.Login)
	a.PUT("/clientes/autogestion/password/:id", ClientesAutogestionController.Password)
	a.PUT("/clientes/autogestion/blanqueo/:id", ClientesAutogestionController.Blanqueo)
	a.GET("/clientes/autogestion/authenticated", ClientesAutogestionController.GetAuthenticatedUser)
	a.PUT("/clientes/autogestion", ClientesAutogestionController.Update)
	a.PUT("/clientes/autogestion/baja/:id", ClientesAutogestionController.Baja)
	a.DELETE("/clientes/autogestion/baja/:id", ClientesAutogestionController.Delete)
	a.PUT("/clientes/autogestion/alta/:id", ClientesAutogestionController.Alta)
	a.GET("/clientes/autogestion/:id", ClientesAutogestionController.Get)
	a.POST("/clientes/autogestion", ClientesAutogestionController.Create)
	a.GET("/clientes/:id/autogestion", ClientesAutogestionController.GetAll)

	//clientes
	a.GET("/clientes/localidades", ClientesDomiciliosController.GetAllLocalidades)
	a.GET("/clientes/:id", ClientesController.Get)
	a.PUT("/clientes", ClientesController.Update)
	a.POST("/clientes", ClientesController.Create)
	a.GET("/clientes", ClientesController.GetAll)
	a.PUT("/clientes/baja/:id", ClientesController.Baja)
	a.PUT("/clientes/alta/:id", ClientesController.Alta)

	//Consulta de obleas
	a.GET("/clientes/:id/obleas", ClientesController.GetObleasAutogestion)
	a.GET("/obleas_compradas/:id", FranquiciasController.GetConObleas)

	//clientes/domicilios
	a.GET("/clientes/domicilios/:id", ClientesDomiciliosController.Get)
	a.PUT("/clientes/domicilios", ClientesDomiciliosController.Update)
	a.POST("/clientes/domicilios", ClientesDomiciliosController.Create)
	a.GET("/clientes/:id/domicilios", ClientesDomiciliosController.GetAll)
	a.DELETE("/clientes/domicilios/:id", ClientesDomiciliosController.Delete)

	//clientes/contactos
	a.GET("/clientes/contactos/:id", ClientesContactosController.Get)
	a.PUT("/clientes/contactos", ClientesContactosController.Update)
	a.POST("/clientes/contactos", ClientesContactosController.Create)
	a.GET("/clientes/:id/contactos", ClientesContactosController.GetAll)
	a.DELETE("/clientes/contactos/:id", ClientesContactosController.Delete)

	//condicioniva
	a.GET("/condicioniva", CondicionivaController.GetAll)
	a.GET("/documentotipos", DocumentotiposController.GetAll)

	//clientes/convenios
	a.GET("/clientes/convenios/:id", ClientesConveniosController.Get)
	a.PUT("/clientes/convenios", ClientesConveniosController.Update)
	a.POST("/clientes/convenios", ClientesConveniosController.Create)
	a.GET("/clientes/:id/convenios", ClientesConveniosController.GetAll)
	a.DELETE("/clientes/convenios/:id", ClientesConveniosController.Delete)

	a.GET("/afip/cuit/:id", AfipController.GetCUIT)

	a.POST("/clientes/resumendecuenta", ClientesController.GetResumendeCuenta)
	a.POST("/clientes/resumendecuentapdf", ClientesController.Imprimir)

	a.POST("/clientes/ctacte/pendientes", ClientesController.GetCtaCtePendienteAutogestion)
	a.POST("/clientes/resumendecuentafranquiciaspdf", ClientesController.ImprimirFAutogestion)

	//clientes/destinos
	a.GET("/clientes/destinos/:id", ClientesDestinosController.Get)
	a.PUT("/clientes/destinos", ClientesDestinosController.Update)
	a.POST("/clientes/destinos", ClientesDestinosController.Create)
	a.GET("/clientes/:id/destinos", ClientesDestinosController.GetAll)
	a.DELETE("/clientes/destinos/:id", ClientesDestinosController.Delete)

	//articulos
	a.GET("/articulos/bultos", ArticulosController.GetBultos)
	a.GET("/articulos/:id", ArticulosController.Get)
	a.PUT("/articulos", ArticulosController.Update)
	a.POST("/articulos", ArticulosController.Create)
	a.GET("/articulos", ArticulosController.GetAll)
	a.DELETE("/articulos/:id", ArticulosController.Delete)
	a.GET("/articulos/get", ArticulosController.GetArticulos)

	a.GET("/provincias", ProvinciasController.GetAll)

	//domiciliotipos
	a.GET("/domicilios_tipos", Domicilios_tiposController.GetAll)

	//servicios
	a.GET("/servicios/:id", ServiciosController.Get)
	a.GET("/servicios", ServiciosController.GetAll)

	a.GET("/autogestion/retiros", RetirosController.GetAllClientes)
	a.POST("/autogestion/retiros", RetirosController.CreateClientes)
	a.GET("/autogestion/retiros/:id", RetirosController.GetClientes)
	a.POST("/autogestion/retiros/destinos", RetirosController.DestinosClientes)
	a.DELETE("/autogestion/retiros/:id", RetirosController.DeleteDestino)
	a.GET("/autogestion/retiros/:id/imprimir", RetirosController.ImprimirClientes)

	//frecuencias
	a.GET("/frecuencias", FrecuenciasController.GetAll)

	a.GET("/track/comprobante/:id", TrackController.ShowTrackingComprobante)
	a.GET("/track/:id", TrackController.ShowTrackCli)
	a.GET("/tracking/:id", TrackController.ShowTrackCli)
	a.GET("/track", Track_lecturasController.GetAllTracks)

	// Rangos horarios
	a.GET("/rangoshorarios", RangosHorariosController.GetAll)
	a.POST("/rangoshorarios", RangosHorariosController.Create)
	a.PUT("/rangoshorarios", RangosHorariosController.Update)

	//Comprobantes tipos
	a.GET("/comprobantestipos", ComprobantesTiposController.GetAll)
	a.GET("/comprobantestipos/:id", ComprobantesTiposController.Get)
	a.POST("/comprobantestipos", ComprobantesTiposController.Create)
	a.PUT("/comprobantestipos", ComprobantesTiposController.Update)

	a.GET("/monedas", EmpresasController.GetMonedas)

	//Comprobantes
	a.POST("/comprobantes", ComprobantesController.Create)
	a.POST("/comprobantes/prepago", ComprobantesController.CreatePrepago)
	a.GET("/comprobantes/:id", ComprobantesController.Get)
	a.DELETE("/comprobantes/:id", ComprobantesController.Anular)
	a.GET("/comprobantes/:id/etiqueta", ComprobantesTiposController.ImprimirEtoqueta)
	a.GET("/comprobantes/:id/imprimir", ComprobantesTiposController.Imprimir)
	a.POST("/comprobantes/:id/enviarmail", ComprobantesTiposController.EnviarporMail)
	a.GET("/comprobantes", ComprobantesController.GetAll)
	a.GET("/comprobantes/destino", ComprobantesController.GetAllDestino)
	a.GET("/comprobantes/obleas", ComprobantesController.GetAllObleas)
	a.GET("/comprobantes/consultaobleas", ComprobantesController.GetAllConsultaObleas)
	a.GET("/comprobantes/:id/etiqueta", ComprobantesTiposController.ImprimirEtoqueta)

	a.GET("/saldos", ComprobantesController.GetSaldo)
	a.GET("/saldos/:id", ComprobantesController.GetSaldoCliente)

	//autogestion
	a.GET("/autogestion/comprobantes", ComprobantesController.GetAllComprobantes)
	a.GET("/autogestion/comprobantes/d", ComprobantesController.GetAllComprobantesDestino)
	a.GET("/autogestion/obleas", ComprobantesController.GetAllObleasClientes)

	a.POST("/iibb/:cuit", ClientesController.GetIibb)

	//entregas
	a.GET("/entregas", EntregasController.GetAll)
	a.GET("/entregas/:id", EntregasController.Get)
	a.GET("/entregas/auto", EntregasController.GetAllAutogestion)

	//coebrtura
	a.POST("/cobertura", FranquiciasLocalidadesController.GetCobertura)
	a.GET("/coberturageo", FranquiciasController.GetCoberturaGeo)
	a.GET("/coberturageoexluir", FranquiciasController.GetCoberturaExcluir)

	a.POST("/autogestion/cobertura", FranquiciasLocalidadesController.GetCoberturaAuto)
	a.POST("/autogestion/remitos/enviar", AutogestionController.Enviar)
	a.POST("/autogestion/remitos/enviarconlote", AutogestionController.EnviarConLote)
	a.GET("/autogestion/:id/imprimir", AutogestionController.Imprimir)
	a.GET("/autogestion/:id/imprimirverificacion", AutogestionController.ImprimirVerificacion)
	a.GET("/autogestion/comprobantes/:id/etiqueta", ComprobantesTiposController.ImprimirEtiquetas)

	//clientes AUTOGESTION agenda
	a.GET("/clientes/autogestion/agenda/:id", AgendaController.Get)
	a.PUT("/clientes/autogestion/agenda/:id", AgendaController.Update)
	a.POST("/clientes/autogestion/agenda", AgendaController.Create)
	a.GET("/clientes/autogestion/:id/agenda", AgendaController.GetAll)
	a.DELETE("/clientes/autogestion/agenda/:id", AgendaController.Delete)

	a.GET("/localidades", LocalidadesController.Getlocalidaddes)
	a.GET("/localidad/:id", LocalidadesController.Get)

	a.POST("/enviar", AutogestionController.Enviar)

	a.POST("/autogestion/enviar/obleas", AutogestionController.CreateOblea)
	a.GET("/constancia/:id/imprimir", ComprobantesTiposController.Constancia)
	a.GET("/autogestion/comprobantes/exportar/:id", ComprobantesController.GetExcel)
	e.Logger.Fatal(e.Start(config.Port))
	// e.Logger.Fatal(e.Start(":8000"))

}
