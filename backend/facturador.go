package main

import (
	"net/http"

	AfipController "./controllers/afip"
	AgilDataController "./controllers/agildata"
	ArticulosController "./controllers/articulos"
	Articulos_distanciasController "./controllers/articulos_distancias"
	Articulos_preciosController "./controllers/articulos_precios"

	AutomaticosController "./controllers/automaticos"
	PromoController "./controllers/promo"

	ClientesAcuerdosController "./controllers/acuerdos"
	ClientesBolsinesController "./controllers/bolsines"
	ClientesController "./controllers/clientes"
	ClientesContactosController "./controllers/clientes_contactos"
	ClientesConveniosController "./controllers/clientes_convenios"
	ClientesDestinosController "./controllers/clientes_destinos"
	ClientesDomiciliosController "./controllers/clientes_domicilios"
	RedespachantesController "./controllers/redespachantes"

	AutComprobantesController "./controllers/aut_comprobantes"
	ComprobantesController "./controllers/comprobantes"

	TrackController "./controllers/track"
	Track_lecturasController "./controllers/track_lecturas"

	BolsinesController "./controllers/bolsines"
	DespachosController "./controllers/despachos"

	CondicionivaController "./controllers/condicioniva"
	DocumentotiposController "./controllers/documentotipos"
	Domicilios_tiposController "./controllers/domicilios_tipos"
	ProvinciasController "./controllers/provincias"

	Accesos_rolesController "./controllers/accesos_roles"
	FranquiciasAcuerdosController "./controllers/acuerdos"
	CajaController "./controllers/caja"
	CajasController "./controllers/cajas"
	CierresController "./controllers/cierres"
	ComprobantesTiposController "./controllers/comprobantes_tipos"
	EmpresasController "./controllers/empresas"
	FranquiciasController "./controllers/franquicias"
	FranquiciasCarterosController "./controllers/franquicias_carteros"
	FranquiciasCoberturaController "./controllers/franquicias_cobertura"
	FranquiciasLocalidadesController "./controllers/franquicias_localidades"
	FranquiciasNodosController "./controllers/franquicias_nodos"
	FranquiciasZonasController "./controllers/franquicias_zonas"
	FrecuenciasController "./controllers/frecuencias"
	LineasController "./controllers/lineas"
	LocalidadesController "./controllers/localidades_ar"
	MaterialesController "./controllers/materiales"
	MaterialesStockController "./controllers/materiales_stock"
	NodosController "./controllers/nodos"
	PuntosDeVentasController "./controllers/puntosdeventas"
	PuntosDeVentasComprobantesController "./controllers/puntosdeventas_comprobantes"
	RangosHorariosController "./controllers/rangos_horarios"
	RecibosController "./controllers/recibos"
	RetirosController "./controllers/retiros"
	RolesController "./controllers/roles"
	RubrosController "./controllers/rubros"
	RuteosController "./controllers/ruteos"
	ServiciosController "./controllers/servicios"
	SistemasController "./controllers/sistemas"
	TroncalesNodosController "./controllers/troncales_nodos"
	UsuariosController "./controllers/usuarios"
	ZonasController "./controllers/zonas"

	ObservacionesController "./controllers/observaciones"
	SincroController "./controllers/sincro"
	Solicitud_ncController "./controllers/solicitud_nc"
	Solicitud_nc_motivosController "./controllers/solicitud_nc_motivos"

	VehiculosController "./controllers/vehiculos"
	Vehiculos_salidasController "./controllers/vehiculos_servicios"

	Articulos_distancias_preciospesoController "./controllers/articulos_distancias_preciospeso"
	Articulos_pesosController "./controllers/articulos_pesos"
	AutorizacionesController "./controllers/autorizaciones"
	BancosController "./controllers/bancos"
	Comprobantes_cobradosController "./controllers/comprobantes_cobrados"
	CotizacionController "./controllers/cotizacion"
	Cotizacion_distanciasController "./controllers/cotizacion_distancias"
	CuentasController "./controllers/cuentas"
	ExigenciasController "./controllers/exigencias"
	Franquicias_articulos_excepcionesController "./controllers/franquicias_articulos_excepciones"
	GuiasController "./controllers/guias"
	ImpuestosController "./controllers/impuestos"
	TarjetasController "./controllers/tarjetas"
	TrabajosController "./controllers/trabajos"

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
	a.GET("/usuarios/cajas", UsuariosController.GetAllParaCajas)

	a.GET("/usuarios/franquicias", FranquiciasController.GetFranquiciasbyusuario)

	a.GET("/usuarios/roles", RolesController.GetRolesbyusuario)

	a.GET("/usuarios/sistemas", SistemasController.GetSistemasbyusuario)

	a.GET("/roles", RolesController.GetAll)
	a.GET("/exigencias", ExigenciasController.GetAll)

	a.PUT("/accesos_roles", Accesos_rolesController.Update)
	a.POST("/accesos_roles", Accesos_rolesController.Create)
	a.GET("/accesos_roles/:id", Accesos_rolesController.GetAll)
	a.DELETE("/accesos_roles/:id/:idacceso", Accesos_rolesController.Delete)

	//usauarios_preferencias
	a.POST("/usuarios/preferencias", UsuariosController.Create_preferens)

	//franquicias
	a.GET("/franquicias/depositos/:id", FranquiciasController.GetDepositos)
	a.GET("/franquicias/all", FranquiciasController.GetAllconNodos)
	a.GET("/franquicias/dominios", FranquiciasController.GetAllDominios)
	a.GET("/franquicias/:ido/:idd", FranquiciasController.GetDistancia)
	a.GET("/franquicias/dashboard", FranquiciasController.GetDashboard)
	a.GET("/franquicias/acuerdos/tipos", FranquiciasAcuerdosController.GetAcuerdos_tipos)
	a.GET("/franquicias/:id/acuerdos", FranquiciasAcuerdosController.GetAllFranquicias)
	a.PUT("/franquicias/acuerdos", FranquiciasAcuerdosController.UpdateFranuicia)
	a.PUT("/franquicias/:id/baja", FranquiciasController.BajaFranquicia)
	a.PUT("/franquicias/:id/alta", FranquiciasController.Alta)
	a.POST("/franquicias/acuerdos", FranquiciasAcuerdosController.Create)
	a.DELETE("/franquicias/acuerdos/:id", FranquiciasAcuerdosController.Baja)

	a.GET("/franquicia/receptorias", FranquiciasController.GetReceptorias)

	a.GET("/franquicias/tipos", FranquiciasController.GetAllTipos)
	a.GET("/franquicias/:id", FranquiciasController.Get)
	a.PUT("/franquicias", FranquiciasController.Update)
	a.POST("/franquicias", FranquiciasController.Create)
	a.GET("/franquicias", FranquiciasController.GetAll)
	a.DELETE("/franquicias/:id", FranquiciasController.Delete)
	a.GET("/franquicias/:id/nodos", FranquiciasController.GetAllNodos)
	a.GET("/franquicias/atiende", FranquiciasController.GetFranquiciaZona)
	a.GET("/franquicias/noatiende", FranquiciasController.GetFranquiciaZonaExcluir)
	a.GET("/franquicias/:id/stock", NodosController.Stock)
	a.GET("/franquicias/acuerdo", FranquiciasController.GetAcuerdo)
	a.GET("/franquicias/remesas", RecibosController.GetAllRemesas)
	a.GET("/franquicias/remesas/sucursales", RecibosController.GetAllRemesassucursales)

	a.GET("/sucursales/remesasctrl", RecibosController.GetAllRemesasCtrlsucusrsales)

	a.GET("/franquicias/remesasctrl", RecibosController.GetAllRemesasCtrl)
	a.GET("/franquicias/remesasctrl/:id", RecibosController.GetRemesaCtrl)
	a.GET("/franquicias/caja", RecibosController.GetAllRemesa)
	a.GET("/franquicias/aRemesar", RecibosController.GetAllaRemesa)
	a.GET("/franquicias/aRemesar/sucursal", RecibosController.GetAllaRemesasucursal)

	a.POST("/franquicias/caja/pdf", RecibosController.GetAllRemesaPdf)
	a.POST("/franquicias/cajaremesa/pdf", RecibosController.GetAllAremesarPdf)
	a.GET("/franquicias/remesa/:id", RecibosController.GetRemesa)
	a.GET("/franquicias/remesa/sucursal/:id", RecibosController.GetRemesasucursal)
	a.GET("/remesa/:id/cheques", RecibosController.GetRemesaCheques)
	a.GET("/remesas/cheques", RecibosController.GetCheques)
	a.GET("/franquicias/red", FranquiciasController.GetAllRed)
	// a.POST("/remesa/:id/gastos", RecibosController.CreateGastos)
	a.POST("/remesa/gastos", RecibosController.CreateGasto)
	a.DELETE("/remesa/gastos/:id", RecibosController.DeleteGasto)
	a.GET("/remesa/gasto/:id", RecibosController.GetGasto)
	a.POST("/remesa/:id/enviar", RecibosController.EnviarRemesa)
	a.DELETE("/remesa/:id/borrar", RecibosController.BorrarRemesa)
	a.DELETE("/remesa/sucursal/:id/borrar", RecibosController.BorrarRemesasucursal)
	a.POST("/remesas_control", RecibosController.Rechazar_Control)
	a.POST("/remesas_control/sucursal", RecibosController.Rechazar_Controlsucursal)
	a.PUT("/remesas/:id", RecibosController.Confirmar)
	a.PUT("/sucursal/remesas/:id", RecibosController.Confirmarsucursal)
	a.POST("/franquicias/remesa/:id/pdf", RecibosController.GetRemesaPdf)
	a.POST("/franquicias/remesa/sucursal/:id/pdf", RecibosController.GetRemesaPdfsucursal)
	a.POST("/franquicias/remesar", RecibosController.Remesar)
	a.POST("/franquicias/remesar/sucursal", RecibosController.Remesarsucursal)
	a.POST("/recibos", RecibosController.Create)
	//franquicias_nodos
	a.GET("/franquicias/nodos/:id", FranquiciasNodosController.Get)
	a.PUT("/franquicias/nodos", FranquiciasNodosController.Update)
	a.POST("/franquicias/nodos", FranquiciasNodosController.Create)
	a.GET("/franquicias/nodos", FranquiciasNodosController.GetAll)
	a.DELETE("/franquicias/nodos/:id", FranquiciasNodosController.Delete)
	//franquicias_zonas
	a.GET("/franquicias/zonas/:id", FranquiciasZonasController.Get)
	a.PUT("/franquicias/zonas", FranquiciasZonasController.Update)
	a.POST("/franquicias/zonas", FranquiciasZonasController.Create)
	a.GET("/franquicias/:id/zonas", FranquiciasZonasController.GetAll)
	a.DELETE("/franquicias/zonas/:id", FranquiciasZonasController.Delete)

	//FranquiciasCobertura
	a.GET("/franquicias/cobertura/:id", FranquiciasCoberturaController.Get)
	a.PUT("/franquicias/cobertura", FranquiciasCoberturaController.Update)
	a.POST("/franquicias/cobertura", FranquiciasCoberturaController.Create)
	a.GET("/franquicias/:id/cobertura", FranquiciasCoberturaController.GetAll)
	a.DELETE("/franquicias/cobertura/:id", FranquiciasCoberturaController.Delete)

	a.POST("/cobertura", FranquiciasLocalidadesController.GetCobertura1)
	a.GET("/coberturageo", FranquiciasController.GetCoberturaGeo)
	a.GET("/coberturageoexluir", FranquiciasController.GetCoberturaExcluir)

	//FranquiciasCobertura
	a.GET("/franquicias/localidades/:id", FranquiciasLocalidadesController.Get)
	a.PUT("/franquicias/localidades", FranquiciasLocalidadesController.Update)
	a.POST("/franquicias/localidades", FranquiciasLocalidadesController.Create)
	a.GET("/franquicias/:id/localidades", FranquiciasLocalidadesController.GetAll)
	a.DELETE("/franquicias/localidades/:id", FranquiciasLocalidadesController.Delete)
	a.GET("/coberturas", FranquiciasLocalidadesController.GetCoberturas)

	//franquicias_carteros
	a.GET("/franquicias_carteros/:id", FranquiciasCarterosController.Get)
	a.PUT("/franquicias_carteros", FranquiciasCarterosController.Update)
	a.PUT("/franquicias_carteros/direccion", FranquiciasCarterosController.UpdateDireccion)
	a.POST("/franquicias_carteros", FranquiciasCarterosController.Create)
	a.GET("/franquicias_carteros", FranquiciasCarterosController.GetAll)
	a.DELETE("/franquicias_carteros/:id", FranquiciasCarterosController.Delete)
	a.GET("/franquicias/usuarios", FranquiciasCarterosController.GetUsuariobyFranquicias)

	a.GET("/clientes/:id/obleas", ClientesController.GetObleas)

	//clientes
	a.GET("/clientes/find", ClientesController.GetAutocompletar)
	a.GET("/clientes/getfind/:id", ClientesController.GetFind)
	a.GET("/clientes/franquicia", ClientesController.GetAllFranquicias)
	a.GET("/clientes/localidades", ClientesDomiciliosController.GetAllLocalidades)
	a.GET("/clientes/:id", ClientesController.Get)
	a.PUT("/clientes", ClientesController.Update)
	a.POST("/clientes", ClientesController.Create)
	a.GET("/clientes", ClientesController.GetAll)

	a.PUT("/clientes/baja/:id", ClientesController.Baja)
	a.PUT("/clientes/alta/:id", ClientesController.Alta)
	a.GET("/clientes_verificar/:cuit", ClientesController.GetVerificar)
	a.GET("/clientes_validacion", ClientesController.Validar)
	a.POST("/clientes_domicilios/:franquicias_id/:id", ClientesDomiciliosController.AsignarSucursal)

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

	//clientes/bolsines
	a.GET("/clientes/bolsines/:id", ClientesBolsinesController.Get)
	a.PUT("/clientes/bolsines", ClientesBolsinesController.Update)
	a.POST("/clientes/bolsines", ClientesBolsinesController.Create)
	a.GET("/clientes/:id/bolsines", ClientesBolsinesController.GetAll)
	a.DELETE("/clientes/bolsines/:id", ClientesBolsinesController.Baja)

	//condicioniva
	a.GET("/condicioniva", CondicionivaController.GetAll)
	a.GET("/condicioniva/:id", CondicionivaController.Get)
	a.GET("/documentotipos", DocumentotiposController.GetAll)

	//clientes/convenios
	a.GET("/clientes/convenios/:id", ClientesConveniosController.Get)
	a.PUT("/clientes/convenios", ClientesConveniosController.Update)
	a.POST("/clientes/convenios", ClientesConveniosController.Create)
	a.GET("/clientes/:id/convenios", ClientesConveniosController.GetAll)
	a.DELETE("/clientes/convenios/:id", ClientesConveniosController.Delete)
	a.GET("/clientes/convenios/:id/franquicias", ClientesConveniosController.GetAllFranquiciasConvenios)
	// a.POST("/clientes/convenios/:id/franquicias", ClientesConveniosController.CreateConveniosFranquicias)

	a.GET("/afip/cuit/:id", AfipController.GetCUIT)
	a.GET("/getByDni", AgilDataController.FetchPersonByDni)
	a.GET("/getByCuit/:cuit", AgilDataController.FetchPersonByCuit)
	//clientes/acuerdos

	a.GET("/clientes/acuerdos/tipos", ClientesAcuerdosController.GetAcuerdos_tipos)
	a.GET("/clientes/acuerdos/:id", ClientesAcuerdosController.Get)
	a.PUT("/clientes/acuerdos", ClientesAcuerdosController.Update)
	a.POST("/clientes/acuerdos", ClientesAcuerdosController.Create)
	a.GET("/clientes/:id/acuerdos", ClientesAcuerdosController.GetAll)
	a.DELETE("/clientes/acuerdos/:id", ClientesAcuerdosController.Baja)
	a.DELETE("/clientes/acuerdos/:id/cancelar", ClientesAcuerdosController.Cancelar)
	a.GET("/clientes/acuerdos/:id/:idarticulo", ClientesAcuerdosController.GetAcuerdosArticulos)

	//clientes/acuerdos
	a.GET("/clientes/bolsines/:id", BolsinesController.Get)
	a.PUT("/clientes/bolsines", BolsinesController.Update)
	a.POST("/clientes/bolsines", BolsinesController.Create)
	a.GET("/clientes/:id/bolsines", BolsinesController.GetAll)
	a.DELETE("/clientes/bolsines/:id", BolsinesController.Delete)

	a.POST("/clientes/resumendecuenta", ClientesController.GetResumendeCuenta)
	a.POST("/clientes/ctacte/pendientes", ClientesController.GetCtaCtePendiente)
	a.POST("/clientes/resumendecuentapdf", ClientesController.Imprimir)
	a.POST("/clientes/resumendecuentafranquiciaspdf", ClientesController.ImprimirF)

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
	a.GET("/articulos/solicitud", ArticulosController.GetAllArticulosSolicitud)
	a.GET("/articulos/lista_precios", ArticulosController.GetAllListaDePrecios)
	a.GET("/articulos/pack", ArticulosController.GetArticulosPack)
	a.GET("/articulos/precios/peso", ArticulosController.GetPreciosPeso)
	a.DELETE("/articulos/:id", ArticulosController.Delete)
	a.PUT("/articulos/baja/:id", ArticulosController.Baja)
	a.GET("/embalajes", ArticulosController.GetAllEmbalajes)
	a.GET("/articulos/get", ArticulosController.GetArticulos)
	a.GET("/articulos/getAdicionales", ArticulosController.GetgetAdicionales)
	a.GET("/articulos/clientes/:id", ArticulosController.GetArticulos_Cli)
	a.GET("/cotizar", ArticulosController.GetCotizar)
	a.GET("/cotizartodos", ArticulosController.GetCotizarTodos)
	a.POST("/articulos/list", ArticulosController.GetArticulosList)
	a.POST("/articulos/list", ArticulosController.GetArticulosList1) //Para nuevo facturador
	a.GET("/articulos_grupos", ArticulosController.GetAllGrupos)

	//articulos_pesos
	a.POST("/articulos_pesos", Articulos_pesosController.Create)
	a.PUT("/articulos_pesos", Articulos_pesosController.Update)
	a.GET("/articulos_pesos", Articulos_pesosController.GetAll)
	a.GET("/articulos_pesos/:id", Articulos_pesosController.Get)
	a.DELETE("/articulos_pesos/:id", Articulos_pesosController.Delete)

	// a.GET("/articulos_peso/get", Articulos_pesosController.GetArticuloPorPeso)

	//articulos_precios
	a.GET("/articulos/precios:id", Articulos_preciosController.Get)
	a.PUT("/articulos/precios", Articulos_preciosController.Update)
	a.POST("/articulos/precios", Articulos_preciosController.Create)
	a.GET("/articulos/:id/precios", Articulos_preciosController.GetAll)
	a.DELETE("/articulos/precios:id", Articulos_preciosController.Delete)

	//articulos_distancias
	a.GET("/articulos_distancias:id", Articulos_distanciasController.Get)
	a.PUT("/articulos_distancias", Articulos_distanciasController.Update)
	a.POST("/articulos_distancias", Articulos_distanciasController.Create)
	// a.GET("/articulos_distancias/:id", Articulos_distanciasController.GetAll)
	a.GET("/articulos_distancias", Articulos_distanciasController.GetAll)
	a.DELETE("/articulos_distancias:id", Articulos_distanciasController.Delete)

	// articulos_distancias_preciospeso
	a.POST("/articulos_distancias_preciospeso", Articulos_distancias_preciospesoController.Create)
	a.PUT("/articulos_distancias_preciospeso", Articulos_distancias_preciospesoController.Update)
	a.GET("/articulos_distancias_preciospeso", Articulos_distancias_preciospesoController.GetAll)
	a.GET("/articulos_distancias_preciospeso/:id", Articulos_distancias_preciospesoController.Get)
	a.DELETE("/articulos_distancias_preciospeso/:id", Articulos_distancias_preciospesoController.Delete)

	//cotizaciones
	a.GET("/cotizacion/:id", CotizacionController.Get)
	a.PUT("/cotizacions", CotizacionController.Update)
	a.POST("/cotizacion", CotizacionController.Create)
	a.POST("/cotizacion/new", CotizacionController.CreateNew)
	a.GET("/cotizacion", CotizacionController.GetAll)
	a.DELETE("/cotizacion/:id", CotizacionController.Delete)
	a.POST("/cotizacion/:id/pdf", CotizacionController.Pdf)
	a.POST("/cotizacion/:id/generar", CotizacionController.Generar)

	//cotizacion_distancias
	a.GET("/cotizacion_distancias/:id", CotizacionController.GetCotizacionDistancias)
	a.GET("/cotizacion_distancias/especifico/:id", CotizacionController.GetCotizacionDistanciasEspecifico)
	a.GET("/cotizacion_distancias/detallado/:id", CotizacionController.GetCotizacionDistanciasDetallado)
	a.POST("/cotizacion_distancias/especifico/editar", Cotizacion_distanciasController.EditarEspecifico)
	a.POST("/cotizacion_distancias/detallado/editar", Cotizacion_distanciasController.EditarDetallado)
	a.POST("/cotizacion_distancias/editar", Cotizacion_distanciasController.Editar)

	a.GET("/cotizacion/autorizaciones/:id", CotizacionController.GetCotizacion)
	a.GET("/cotizacion/autorizaciones/especifico", CotizacionController.GetCotizacionAutorizacionesEspecifico)
	a.GET("/cotizacion/autorizaciones/especifico/pack", CotizacionController.GetCotizacionAutorizacionesEspecificoPack)
	a.GET("/cotizacion/autorizaciones/especifico/rangos", CotizacionController.GetCotizacionAutorizacionesEspecificoRangos)
	a.GET("/cotizacion/autorizaciones/especifico/rangos/pack", CotizacionController.GetCotizacionAutorizacionesEspecificoRangosPack)
	a.GET("/cotizacion/autorizaciones/detallado/:id", CotizacionController.GetCotizacionAutorizacionesDetallado)
	a.GET("/cotizacion/:id/autorizaciones", CotizacionController.GetCotizacionAutorizacion)
	a.POST("/cotizacion/autorizaciones/acuerdo/especifico", CotizacionController.GenerarAcuerdoEspecifico)
	a.POST("/cotizacion/autorizaciones/acuerdo/detallado", CotizacionController.GenerarAcuerdoDetallado)
	a.POST("/cotizacion/autorizaciones/acuerdo", CotizacionController.GenerarAcuerdo)
	a.POST("/cotizacion/:id/imprimir", CotizacionController.Imprimir)

	// Acuerdos
	a.GET("/acuerdos/cotizacion/:id", ClientesAcuerdosController.GetAcuerdo)
	a.GET("/acuerdos/cotizacion/especifico", ClientesAcuerdosController.GetAcuerdos_articulos_distanciaEspecifico)
	a.GET("/acuerdos/cotizacion/especifico/pack", ClientesAcuerdosController.GetAcuerdos_articulos_distanciaEspecificoPack)
	a.GET("/acuerdos/cotizacion/especifico/rangos", ClientesAcuerdosController.GetAcuerdos_articulos_distanciaEspecificoRangos)
	a.GET("/acuerdos/cotizacion/especifico/rangos/pack", ClientesAcuerdosController.GetAcuerdos_articulos_distanciaEspecificoRangosPack)
	a.GET("/acuerdos/cotizacion/detallado/:id", ClientesAcuerdosController.GetAcuerdos_articulos_distanciaDetallado)

	a.GET("/acuerdos/:id/cotizacion", ClientesAcuerdosController.GetAcuerdos_articulos_distancia)
	a.GET("/acuerdos/cotizacion", ClientesAcuerdosController.GetAllAcuerdos)

	a.POST("/acuerdos/cotizacion/especifico/editar", ClientesAcuerdosController.EditarEspecifico)
	a.POST("/acuerdos/cotizacion/detallado/editar", ClientesAcuerdosController.EditarDetallado)
	a.POST("/acuerdos/cotizacion/editar", ClientesAcuerdosController.Editar)

	a.POST("/acuerdos/cotizacion/:id/especifico/imprimir", ClientesAcuerdosController.ImprimirAcuerdoEspecifico)
	a.POST("/acuerdos/cotizacion/:id/detallado/imprimir", ClientesAcuerdosController.ImprimirAcuerdoDetallado)
	a.POST("/acuerdos/cotizacion/:id/general/imprimir", ClientesAcuerdosController.ImprimirAcuerdoGeneral)
	a.POST("/acuerdos/cotizacion/:id/valordeclarado/imprimir", ClientesAcuerdosController.ImprimirAcuerdoValorDeclarado)

	//Fin Rutas
	a.GET("/provincias/:id", ProvinciasController.Get)
	a.GET("/provincias", ProvinciasController.GetAll)

	//domiciliotipos
	a.GET("/domicilios_tipos/:id", Domicilios_tiposController.Get)
	a.PUT("/domicilios_tipos", Domicilios_tiposController.Update)
	a.POST("/domicilios_tipos", Domicilios_tiposController.Create)
	a.GET("/domicilios_tipos", Domicilios_tiposController.GetAll)

	//rubros
	a.GET("/rubros/:id", RubrosController.Get)
	a.PUT("/rubros", RubrosController.Update)
	a.POST("/rubros", RubrosController.Create)
	a.GET("/rubros", RubrosController.GetAll)
	a.DELETE("/rubros/:id", RubrosController.Delete)

	//servicios
	a.GET("/servicios/:id", ServiciosController.Get)
	a.PUT("/servicios", ServiciosController.Update)
	a.POST("/servicios", ServiciosController.Create)
	a.GET("/servicios", ServiciosController.GetAll)
	a.DELETE("/servicios/:id", ServiciosController.Delete)
	//servicios
	a.GET("/zonas/:id", ZonasController.Get)
	a.PUT("/zonas", ZonasController.Update)
	a.POST("/zonas", ZonasController.Create)
	a.GET("/zonas", ZonasController.GetAll)
	a.DELETE("/zonas/:id", ZonasController.Delete)

	//nodos
	//	a.GET("/nodos/:id/ingreso/franquicia/:idfranquicia", NodosController.IngresoFranquicia)

	a.GET("/nodos/dominios", NodosController.Getdominios)
	a.GET("/nodos/:id/dominios", NodosController.Get_dominios)

	a.GET("/nodos/:id/auditoria", NodosController.Auditar)
	//a.POST("/nodos/auditoria", NodosController.Auditoria)

	a.GET("/nodos/:id/cargar/:dominio", NodosController.Cargar)
	a.GET("/nodos/:id/descargar/:dominio", NodosController.Descargar)

	a.GET("/nodos/:id/transbordar/:dominio/:dominiod", NodosController.Transbordar)
	a.POST("/nodos/transbordar", NodosController.TransbordarCamion)

	a.POST("/nodos/cargar", NodosController.CargarCamion)
	a.POST("/nodos/descargar", NodosController.DescargarCamion)

	a.GET("/nodos/:nodo/ingreso/:id", NodosController.Ingreso)
	a.POST("/nodos/ingreso", NodosController.IngresoFranquicia)

	a.GET("/nodos/:nodo/egreso/:id", NodosController.Egreso)
	a.POST("/nodos/egresos", NodosController.EgresoFranquicia)

	a.POST("/nodos/rutas/cfg", NodosController.Addnodo)
	a.DELETE("/nodos/rutas/cfg", NodosController.Delnodo)
	a.GET("/nodos/rutas", NodosController.GetAllNodosRutas)
	a.GET("/nodos/rutas/:id", NodosController.GetAllNodosRuta)
	a.GET("/nodos/rutas/:id/:lineas_id", NodosController.GetAllNodosRutaCfg)

	a.GET("/nodos/stock", NodosController.StockDeposito)
	a.GET("/nodos/stock/sucursal", NodosController.StockSucursal)
	a.GET("/nodos/stock/sucursal", NodosController.StockSucursal)
	a.GET("/nodos/stock/sucursal/:id", NodosController.StockSucursales)
	a.GET("/nodos/stock/sucursaltransito/:id", NodosController.StockSucursalesTransito)

	a.GET("/nodos/stock/destino", NodosController.StockDestino)
	a.GET("/nodos/todos", NodosController.GetAllNodos)
	a.GET("/nodos/todos/red", NodosController.GetAllNodosRed)
	a.GET("/nodos/tipos", FranquiciasController.GetAllTiposNodos)
	a.GET("/nodos/:id", NodosController.Get)
	a.PUT("/nodos", NodosController.Update)
	a.POST("/nodos", NodosController.Create)
	a.GET("/nodos", NodosController.GetAll)
	a.GET("/nodos/franquicias/:id", NodosController.GetAllFranquicias)
	a.DELETE("/nodos/:id", NodosController.Delete)

	a.GET("/nodos/:id/stock", NodosController.Stock)

	//lineas

	a.GET("/lineas/stock", LineasController.StockTransito)
	a.GET("/lineas/:id", LineasController.Get)
	a.PUT("/lineas", LineasController.Update)
	a.POST("/lineas", LineasController.Create)
	a.GET("/lineas", LineasController.GetAll)
	a.DELETE("/lineas/:id", LineasController.Delete)
	a.GET("/lineas/:id/stock", LineasController.Stock)

	a.GET("/lineas/dominios", LineasController.GetAllDominios)
	a.GET("/lineas/stocktransitoresumen", LineasController.StockTransitoResumen)
	a.GET("/lineas/:id/stocktransitodetalle", LineasController.StockTransitoDetalle)
	a.GET("/lineas/:id/imprimir", LineasController.Imprimir)

	//clientes/domicilios
	a.GET("/lineas/all", TroncalesNodosController.GetAllLineas)
	a.GET("/lineas/nodos/:id", TroncalesNodosController.Get)
	a.PUT("/lineas/nodos", TroncalesNodosController.Update)
	a.POST("/lineas/nodos", TroncalesNodosController.Create)
	a.GET("/lineas/:id/nodos", TroncalesNodosController.GetAll)
	a.DELETE("/lineas/nodos/:id", TroncalesNodosController.Delete)

	//retiros
	a.GET("/retiros/count", RetirosController.GetAllCount)
	a.GET("/retiros/:id", RetirosController.Get)
	a.GET("/retiros/:id/afacturar", RetirosController.GetaFacturar)
	a.POST("/retiros", RetirosController.Create)
	a.GET("/retiros", RetirosController.GetAll)
	a.GET("/retiros/p", RetirosController.GetAllPendientes)
	a.PUT("/retiros/:id/cancelar", RetirosController.Cancelar)
	a.PUT("/retiros/:id/rechazar", RetirosController.Rechazar)
	a.PUT("/retiros/:id/aprobar", RetirosController.Aprobar)
	a.GET("/retiros/:id/imprimir", RetirosController.Imprimir)
	a.GET("/retiros/imprimir/todosp", RetirosController.ImprimirCompsP)
	a.GET("/retiros/imprimir/todosA", RetirosController.ImprimirCompsA)
	a.GET("/retiros/imprimirfactura/:id", RetirosController.ImprimirDetalle)

	//ruteos
	a.GET("/recorrido", RuteosController.GetRecorrido)
	a.POST("/recorrido", RuteosController.GetRecorrido)
	a.POST("/ruteos/all", RuteosController.CreateAll)
	a.POST("/ruteos/zonificar", RuteosController.Zonificar)
	a.POST("/ruteos/nuevo", RuteosController.CreateRuteo)
	a.POST("/ruteos", RuteosController.Create)
	a.GET("/ruteos/:id", RuteosController.Get)
	a.DELETE("/ruteos/:id", RuteosController.Delete)
	a.POST("/ruteos/:id/imprimir", RuteosController.Imprimir)
	a.POST("/ruteos/:id/imprimircobrado", RuteosController.ImprimirCobrado)
	a.PUT("/ruteos/:id/enviara/:idcartero", RuteosController.Enviara)

	a.GET("/ruteos", RuteosController.GetAll)
	a.GET("/ruteos/:id/guias", RuteosController.GetAllGuias)
	a.PUT("/ruteos/guias", RuteosController.SetAllGuias)
	a.PUT("/ruteos/:id/salida", RuteosController.SetSalida)
	a.PUT("/ruteos/:id/llegada", RuteosController.SetLlegada)
	a.PUT("/reordenar", RuteosController.SetReordenar)
	a.GET("/distribucion/mostrador", RuteosController.GetAllMostrador)
	a.GET("/distribucion", RuteosController.GetAllPendientes)
	a.POST("/distribucion/imprimir", RuteosController.ImprimirDistribucion)
	a.POST("/imprimirdistribucion", RuteosController.PrintDistribucion)
	a.POST("/distribucion/mostrador/imprimir", RuteosController.PrintDistribucionMostrador)
	a.GET("/distribucion", RuteosController.GetAllPendientes)
	a.GET("/distribucion/:id", RuteosController.GetbytrackMostrador)
	a.GET("/distribucion/:ruteos_id/:id", RuteosController.GetbytrackRuta)
	a.PUT("/distribucion/enviara/:idcartero", RuteosController.DistribucionEnviara)
	a.PUT("/distribucion/ubicar", RuteosController.SetUbicar)
	a.PUT("/distribucion/unmostrador", RuteosController.Unmostrador)

	//a.POST("/cartero", RuteosController.GetCartero)
	a.DELETE("/ruteos/guias/:id", RuteosController.Borrar)
	a.PUT("/ruteos/entrega", RuteosController.Entregar)

	//materiales
	a.GET("/materiales", MaterialesController.GetAll)
	//materiales_stock
	a.GET("/materiales/stock/:id", MaterialesStockController.Get)
	a.POST("/materiales/stock", MaterialesStockController.Create)
	a.GET("/materiales/:idfranquicia/stock/:id", MaterialesStockController.GetAll)
	a.GET("/materiales/stock/:id/sucursales", MaterialesStockController.GetAllFranquicias)
	a.GET("/materiales/:id/stock", MaterialesStockController.GetConsulta)
	a.PUT("/materiales/stock", MaterialesStockController.Update)
	a.GET("/materiales/stock", MaterialesStockController.GetAllMaterialesInter)
	a.DELETE("/materiales/stock/:id", MaterialesStockController.Delete)

	//frecuencias
	a.GET("/frecuencias", FrecuenciasController.GetAll)

	a.GET("/redespachantes", RedespachantesController.GetAll)
	a.GET("/redespachantes/:id", RedespachantesController.Get)

	//Hoja de ruta
	a.GET("/redespachar/:id", DespachosController.Redespachar)
	a.GET("/hojaderuta", DespachosController.GetAllHojadeRuta)
	a.POST("/hojaderuta", DespachosController.Create)
	a.POST("/hojaderuta/generar", DespachosController.GenerarBolsines)
	a.POST("/hojaderuta/bolsa", DespachosController.CreateBolsa)
	a.POST("/hojaderuta/setbolsa", DespachosController.SetBolsa)
	a.POST("/hojaderuta/setcrearbolsa", DespachosController.SetCrearBolsa)

	a.POST("/hojaderuta/cerrarbolsa", DespachosController.CerrarBolsa)

	a.GET("/hojaderuta/bolsa/:id/imprimir", DespachosController.ImprimirBolsa)
	a.GET("/hojaderuta/bolsa/:id/imprimir/PDF", DespachosController.ImprimirBolsa)
	a.GET("/hojaderuta/bolsa/:id/imprimir/EPL", DespachosController.ImprimirBolsaEPL)
	a.GET("/hojaderuta/bolsa/:id/imprimir/ZPL", DespachosController.ImprimirBolsaZPL)
	a.POST("/hojaderuta/bolsa/descartar", DespachosController.DescartrBolsa)
	a.POST("/hojaderuta/bolsa/cerrar", DespachosController.CerrarBolsa)
	a.GET("/despachos", DespachosController.GetAll)
	a.GET("/despachos/:id", DespachosController.Get)
	a.GET("/despachos/:id/imprimir", DespachosController.ImprimirDespacho)
	a.GET("/despachos/:id/informe", DespachosController.InformeDespacho)
	a.GET("/bolsines/imprimir", DespachosController.ImprimirBolsines)
	a.DELETE("/despachos/comprobante/:id", DespachosController.Quitar)

	a.GET("/track/comprobante/:id", TrackController.ShowTrackingComprobante)
	a.GET("/track/:id", TrackController.ShowTracking)
	a.GET("/tracking/:id", TrackController.ShowTrackCli)
	a.GET("/track", Track_lecturasController.GetAllTracks)

	//Lecturas
	a.GET("/track_lecturas", Track_lecturasController.GetAll)
	a.GET("/track_lecturas/:id", Track_lecturasController.Get)
	a.DELETE("/track_lecturas/:id", Track_lecturasController.Delete)
	a.GET("/track_lecturas/:id/imprimir", Track_lecturasController.Imprimir)
	a.PUT("/track_lecturas/:id/confirmar", Track_lecturasController.Confirmar)
	a.POST("/track_lecturas/ingresonodoresumen", Track_lecturasController.GetIngresoNodoResumen)
	a.POST("/track_lecturas/ingresonododetalle", Track_lecturasController.PdfIngresoNodoDetalle)
	a.POST("/track_lecturas/entregasucursalresumen", Track_lecturasController.GetEntregaSucursalResumen)
	a.POST("/track_lecturas/entregaSucursaldetalle", Track_lecturasController.PdfEntregaSucursalDetalle)
	a.POST("/track_lecturas/cargaresumen", Track_lecturasController.GetCargaResumen)
	a.POST("/track_lecturas/cargadetalle", Track_lecturasController.PdfCargaDetalle)
	a.POST("/track_lecturas/descargaresumen", Track_lecturasController.GetDescargaResumen)
	a.POST("/track_lecturas/descargadetalle", Track_lecturasController.PdfDescargaDetalle)

	a.GET("/rangoshorarios", RangosHorariosController.GetAll)
	a.POST("/rangoshorarios", RangosHorariosController.Create)
	a.PUT("/rangoshorarios", RangosHorariosController.Update)

	a.GET("/rangoshorarios/:id", RangosHorariosController.Get)
	a.DELETE("/rangoshorarios/:id", RangosHorariosController.Delete)

	//Puntos de ventas
	a.GET("/puntosdeventas", PuntosDeVentasController.GetAll)
	a.GET("/puntosdeventas/:id", PuntosDeVentasController.Get)
	a.POST("/puntosdeventas", PuntosDeVentasController.Create)
	a.PUT("/puntosdeventas", PuntosDeVentasController.Update)
	a.DELETE("/puntosdeventas/:id", PuntosDeVentasController.Delete)

	//Comprobantes tipos
	a.GET("/comprobantestipos", ComprobantesTiposController.GetAll)
	a.GET("/comprobantestipos/:id", ComprobantesTiposController.Get)
	a.POST("/comprobantestipos", ComprobantesTiposController.Create)
	a.PUT("/comprobantestipos", ComprobantesTiposController.Update)
	a.DELETE("/comprobantestipos/:id", ComprobantesTiposController.Delete)

	//Empresas
	a.GET("/empresas", EmpresasController.GetAll)
	a.GET("/empresas/:id", EmpresasController.Get)
	a.POST("/empresas", EmpresasController.Create)
	a.PUT("/empresas", EmpresasController.Update)
	a.DELETE("/empresas/:id", EmpresasController.Delete)

	a.GET("/monedas", EmpresasController.GetMonedas)

	//Puntos de ventas comprobantes
	a.GET("/puntosdeventascomprobantes", PuntosDeVentasComprobantesController.GetAll)
	a.GET("/puntosdeventascomprobantes/:id", PuntosDeVentasComprobantesController.Get)
	a.POST("/puntosdeventascomprobantes", PuntosDeVentasComprobantesController.Create)
	a.PUT("/puntosdeventascomprobantes", PuntosDeVentasComprobantesController.Update)
	a.DELETE("/puntosdeventascomprobantes/:id", PuntosDeVentasComprobantesController.Delete)

	//Comprobantes
	a.POST("/comprobantes/redireccionar", ComprobantesController.Redireccionar)
	a.POST("/comprobantes/RedireccionarObleas", ComprobantesController.RedireccionarObleas)
	a.POST("/comprobantes/RedireccionarBultos", ComprobantesController.RedireccionaryRefacturar)

	a.POST("/comprobantes/intermateriales", ComprobantesController.Intermateriales)
	a.POST("/comprobantes/intersucursales", ComprobantesController.Intersucursales)

	a.POST("/comprobantes/facturacion", ClientesController.Facturacion)
	a.POST("/comprobantes/facturaciondestino", ClientesController.FacturacionDestino)
	a.POST("/comprobantes/facturacion/imprimir", ClientesController.ImprimirFacturacion)
	a.POST("/comprobantes/facturacion/imprimirfacturaciondestino", ClientesController.ImprimirFacturacionDestino)
	a.POST("/comprobantes/facturacion/caja/imprimir", ClientesController.ImprimirFacturacionDiaria)
	a.POST("/comprobantes/facturacion/caja/imprimirverificacion", ClientesController.ImprimirFacturacionDiariaVerificacion)
	a.GET("/comprobantes/:id/entregas", ComprobantesController.GetEntregas)

	a.PUT("/comprobantes/:id/precinto/:precinto", ComprobantesController.Precinto)
	a.POST("/comprobante", ComprobantesController.CreateComprobante)
	a.POST("/comprobante/nc", ComprobantesController.CreateComprobanteNC)
	a.POST("/comprobantes", ComprobantesController.Create)
	a.POST("/comprobantes/prepago", ComprobantesController.CreatePrepago)
	a.POST("/comprobantes/nc_parcial", ComprobantesController.CreateNcParcial)
	a.POST("/comprobantes/nc_siniestro", ComprobantesController.CreateNcSiniestro)
	a.POST("/comprobantes/factura", ComprobantesController.CreateFactura)
	a.POST("/comprobantes/nc_saldo", ComprobantesController.CreateNcSaldo)
	a.GET("/comprobantes/nc", ComprobantesController.GetAllNc)
	a.POST("/comprobantes/nc/:id", ComprobantesController.CreateNC)
	a.GET("/comprobantes/:id", ComprobantesController.Get)
	a.DELETE("/comprobantes/:id", ComprobantesController.Anular)
	a.GET("/comprobantes/:id/etiqueta", ComprobantesTiposController.ImprimirEtoqueta)
	a.GET("/comprobantes/:id/etiqueta/EPL", ComprobantesTiposController.ImprimirEtiquetaEPL)
	a.GET("/comprobantes/:id/etiqueta/ZPL", ComprobantesTiposController.ImprimirEtiquetaZPL)
	a.GET("/comprobantes/:id/etiqueta/PDF", ComprobantesTiposController.ImprimirEtoqueta)

	a.GET("/comprobantes/:id/etiqueta_canasto/:idcanasto", ComprobantesTiposController.ImprimirEtiquetaCanasto)

	a.GET("/comprobantes/:id/imprimir", ComprobantesTiposController.Imprimir1)
	// a.GET("/comprobantes/:id/constancia/imprimir", ComprobantesTiposController.ImprimirConstancia)
	a.GET("/comprobantes/:id/remitos", ComprobantesTiposController.ImprimirRemito)
	a.POST("/comprobantes/:id/enviarmail", ComprobantesTiposController.EnviarporMail)
	a.GET("/comprobantes/numero/:id", ComprobantesController.Get_comprobante1)
	a.GET("/comprobantes", ComprobantesController.GetAll)
	a.GET("/comprobantes/destino", ComprobantesController.GetAllDestino)
	a.GET("/comprobantes/obleas", ComprobantesController.GetAllObleas)
	a.GET("/comprobantes/consultaobleas", ComprobantesController.GetAllConsultaObleas)
	a.GET("/obleas/:id/imprimir", ComprobantesTiposController.ImprimirObleasFacturadas)

	a.GET("/aut_comprobantes/autorizacion/:id", AutComprobantesController.GetPorAutorizacion)

	a.GET("/remitos/:id", ComprobantesController.GetRemitosCliente)
	a.GET("/remitos/pendientes", ComprobantesController.GetAllRemitos)
	a.POST("/remitos/facturar", ComprobantesController.Facturar)
	a.GET("/guias", GuiasController.GetAll)

	a.POST("/solicitud_nc_motivos", Solicitud_nc_motivosController.Create)
	a.GET("/solicitud_nc_motivos/:id", Solicitud_nc_motivosController.Get)
	a.GET("/solicitud_nc_motivos", Solicitud_nc_motivosController.GetAll)
	a.PUT("/solicitud_nc_motivos", Solicitud_nc_motivosController.Update)
	a.DELETE("/solicitud_nc_motivos/:id", Solicitud_nc_motivosController.Delete)
	a.PUT("/solicitud_nc_motivos/baja/:id", Solicitud_nc_motivosController.Baja)
	a.PUT("/solicitud_nc_motivos/alta/:id", Solicitud_nc_motivosController.Alta)

	a.POST("/solicitud_nc", Solicitud_ncController.Create)
	a.GET("/solicitud_nc/:id", Solicitud_ncController.Get)
	a.GET("/solicitud_nc/autorizacion/:id", Solicitud_ncController.GetParaAutorizacion)
	a.GET("/solicitud_nc", Solicitud_ncController.GetAll)
	a.DELETE("/solicitud_nc/:id", Solicitud_ncController.Delete)
	a.PUT("/solicitud_nc/generar/:id", Solicitud_ncController.Update)
	a.PUT("/solicitud_nc/baja/:id", Solicitud_ncController.Baja)
	a.PUT("/solicitud_nc/alta/:id", Solicitud_ncController.Alta)

	a.POST("/devolucion/:id", ComprobantesController.DevolucionSimple) //Funciona para ambas devoluciones

	a.POST("/autorizaciones", AutorizacionesController.Create)
	a.GET("/autorizaciones/:id", AutorizacionesController.Get)
	a.GET("/autorizaciones", AutorizacionesController.GetAll)
	a.GET("/autorizaciones/cotizacion", AutorizacionesController.GetAllAutorizaciones)
	a.GET("/autorizaciones/alertas", AutorizacionesController.GetAllAlertas)
	a.POST("/autorizaciones/acuerdos/rechazar", AutorizacionesController.RechazarAcuerdo)
	a.GET("/autorizaciones/solicitud_nc/:id", AutorizacionesController.GetSolicitud)
	a.POST("/autorizaciones/:id/solicitud", AutorizacionesController.SolicitarAutorizacion)

	a.PUT("/autorizaciones/aprobar/:id", AutorizacionesController.Aprobar)
	a.PUT("/autorizaciones/rechazar/:id", AutorizacionesController.Rechazar)
	a.POST("/autorizaciones/autorizar_stop/:id", AutorizacionesController.AprobarStop)
	a.POST("/autorizaciones/rechazar_stop/:id", AutorizacionesController.RechazarStop)
	a.POST("/autorizaciones/autorizar_facturas/:id", AutorizacionesController.AprobarFacturas)
	a.POST("/autorizaciones/rechazar_facturas/:id", AutorizacionesController.RechazarFacturas)

	//Sincronizaciones
	a.POST("/sincro/config", SincroController.GetConfigurar)
	a.POST("/sincro/usuarios", SincroController.GetUsuarios)

	a.GET("/sincro/lineas", SincroController.GetLineas_sincro)
	a.GET("/sincro/obs", SincroController.GetObs)
	a.POST("/sincro/franquicias", SincroController.GetFranquicias)
	a.GET("/sincro/nodos", SincroController.GetNodos_sincro)
	a.POST("/sincro/articulos", SincroController.GetArticulos)
	a.POST("/sincro/clientes", SincroController.GetClientes)
	a.POST("/sincro/clientes/new", ClientesController.Sincro)
	a.POST("/sincro/scripts", SincroController.GetScripts)
	a.POST("/sincro/localidades", SincroController.GetLocalidades)

	a.GET("/obleas/:id", ComprobantesController.GetOblea)
	a.GET("/obleascliente/:id", ComprobantesController.GetObleasCliente)

	//observaciones
	a.POST("/observaciones", ObservacionesController.Create)
	a.PUT("/observaciones", ObservacionesController.Update)
	a.GET("/observaciones", ObservacionesController.GetAll)
	a.GET("/observaciones/:id", ObservacionesController.Get)
	a.DELETE("/observaciones/:id", ObservacionesController.Delete)

	//vehiculos
	a.POST("/vehiculos", VehiculosController.Create)
	a.PUT("/vehiculos", VehiculosController.Update)
	a.GET("/vehiculos", VehiculosController.GetAll)
	a.GET("/vehiculos/disponibles", VehiculosController.GetAllDisponibles)
	a.GET("/vehiculos/:id", VehiculosController.Get)
	a.DELETE("/vehiculos/:id", VehiculosController.Delete)

	a.GET("/vehiculos/salidas", Vehiculos_salidasController.GetAll)
	a.POST("/vehiculos/salidas", Vehiculos_salidasController.Create)
	a.POST("/vehiculos/vuelta/:id", Vehiculos_salidasController.Vuelta)
	a.DELETE("/vehiculos/salidas/:id", Vehiculos_salidasController.Delete)
	a.PUT("/vehiculos/llegada/:id", Vehiculos_salidasController.Update)
	//vehiculos_lineas

	a.GET("/bancos", BancosController.GetAll)

	a.GET("/tarjetas", TarjetasController.GetAll)

	a.POST("/trabajos", TrabajosController.Create)
	a.GET("/trabajos/:id", TrabajosController.Get)
	a.GET("/trabajos", TrabajosController.GetAll)
	a.PUT("/trabajos", TrabajosController.Update)

	a.POST("/iibb/:cuit", ClientesController.GetIibb)

	//cuentas
	a.POST("/cuentas", CuentasController.Create)
	a.PUT("/cuentas", CuentasController.Update)
	a.GET("/cuentas", CuentasController.GetAll)
	a.GET("/cuentas/:id", CuentasController.Get)
	a.DELETE("/cuentas/:id", CuentasController.Delete)

	//franquicias_articulos_excepciones
	a.POST("/franquicias_articulos_excepciones", Franquicias_articulos_excepcionesController.Create)
	a.PUT("/franquicias_articulos_excepciones", Franquicias_articulos_excepcionesController.Update)
	a.GET("/franquicias_articulos_excepciones", Franquicias_articulos_excepcionesController.GetAll)
	a.GET("/franquicias_articulos_excepciones/:id", Franquicias_articulos_excepcionesController.Get)
	a.DELETE("/franquicias_articulos_excepciones/:id", Franquicias_articulos_excepcionesController.Delete)

	//impuestos
	a.POST("/comprobantes/impuestos", ComprobantesController.GetImpuestos)
	a.POST("/impuestos", ImpuestosController.Create)
	a.PUT("/impuestos", ImpuestosController.Update)
	a.GET("/impuestos", ImpuestosController.GetAll)
	a.GET("/impuestos/:id", ImpuestosController.Get)
	a.DELETE("/impuestos/:id", ImpuestosController.Delete)

	//CAJAS
	a.GET("/cajas/:id", CajasController.Get)
	a.PUT("/cajas", CajasController.Update)
	a.PUT("/cajas/baja/:id", CajasController.Baja)
	a.POST("/cajas", CajasController.Create)
	a.GET("/cajas", CajasController.GetAll)
	a.GET("/cajas_franquicia", CajasController.GetAllCajaFranq)
	a.DELETE("/cajas/:id", CajasController.Delete)

	//CAJA

	a.POST("/caja/cerrar", CierresController.Create)
	a.GET("/caja/:id", CajaController.Get)
	a.PUT("/caja", CajaController.Update)
	a.POST("/caja", CajaController.Create)
	a.POST("/caja/transferencias", RecibosController.TransferirCaja)
	a.POST("/caja/ingreso", CajaController.Ingreso)
	a.POST("/caja/egreso", CajaController.Egreso)
	a.GET("/caja/arqueoPDF", CajaController.Arqueocaja)
	// a.GET("/caja/egreso", CajaController.EgresoPDF)
	a.GET("/caja", CajaController.GetAll)
	a.DELETE("/caja/:id", CajaController.Delete)
	a.DELETE("/caja/ajustar", CajaController.Ajustar)

	a.GET("/colven", ComprobantesTiposController.Colven)

	a.GET("/comprobantes_cobrados/resumen", Comprobantes_cobradosController.GetAllResumen)
	a.PUT("/comprobantes_cobrados", Comprobantes_cobradosController.Update)
	a.GET("/comprobantes_cobrados", Comprobantes_cobradosController.GetAll)

	a.GET("/localidades", LocalidadesController.Getlocalidaddes)
	a.GET("/localidad/:id", LocalidadesController.Get)

	//Localidades
	a.POST("/localidades_ar", LocalidadesController.Create)
	a.PUT("/localidades_ar", LocalidadesController.Update)
	a.GET("/localidades_ar", LocalidadesController.GetAll)
	a.GET("/localidades_ar/:id", LocalidadesController.Get)
	a.GET("/localidades_ar/:id/new", LocalidadesController.GetLocalidades_ar)
	a.DELETE("/localidades_ar/:id", LocalidadesController.Delete)

	a.POST("/promo", PromoController.Create)
	a.GET("/promo/:id", PromoController.Get)
	a.GET("/promo", PromoController.GetAll)

	a.GET("/comprobantes_automaticos", AutomaticosController.GetAll)
	a.POST("/stopdebit/:id", AutomaticosController.Stop)
	a.GET("/comprobantes_automaticos/comprobante/:id", AutomaticosController.GetPorComprobante)
	a.GET("/proximolunes", AutomaticosController.GetLunes)
	a.GET("/facturacion/autorizaciones", AutorizacionesController.GetAllAutorizacionesFacturador)
	e.Logger.Fatal(e.Start(config.PortFacturador))
	//e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
