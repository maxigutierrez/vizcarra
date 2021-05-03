package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SqlString      string `json:"SqlString"`
	SqlString_t    string `json:"SqlString_t"`
	Port           string `json:"Port"`
	PortSincro     string `json:"PortSincro"`
	PortDeposito string `json:"PortDeposito"`
	PortLiq        string `json:"PortLiq"`
	UrlStatic      string `json:"UrlStatic"`
	PortFacturador string `json:"PortFacturador"`
	LogMode        bool   `json:"LogMode"`
}

var Conf Config
var SqlString string
var SqlString_t string
var Port string
var PortFacturador string
var PortSincro string
var PortDeposito string
var PortLiq string
var LogMode bool

const Estados_retiros_Pendiente = 1
const Estados_retiros_Aceptado = 2
const Estados_retiros_Rechazado = 3
const Estados_retiros_Facturado = 4
const Estados_retiros_Cancelado = 5
const Estados_retiros_Retirado = 6
const Estados_retiros_Borrador = 7
const Estados_retiros_Noretirado = 8

const UserAdmin = 1
const SucursalCasaCentral = 8     //Casa Central
const PuntodeVentaPrepago = 10443 //Casa Central
const PuntodeVentaRemitos = 2896
const PuntodeVentaFacturacionRemitos = 10443
const PuntodeVentaComprobanteRecibos = 2015
const PuntodeVentaComprobanteNCA = 3070
const PuntodeVentaComprobanteNCB = 3071
const PuntodeVentaComprobanteFA = 2013
const PuntodeVentaComprobanteFB = 2014
const PuntodeVentaComprobanteAD = 3078
const PuntodeVentaComprobanteBolsines = 1
const ArticuloFacturaRemito = 4423
const ComprobantesTipoRecibo = 5
const Prepago = 8
const PuntodeVentaComprobanteINTER = 3146

const UrlStatic = "public/entregas/"
const UrlSocket = "10.8.1.1:3010"
func Init() {
	Conf = LoadConfiguration("config.ini")
	SqlString = Conf.SqlString
	SqlString_t = Conf.SqlString_t
	LogMode = Conf.LogMode
	Port = Conf.Port
	PortSincro = Conf.PortSincro
	PortLiq = Conf.PortLiq
	PortFacturador = Conf.PortFacturador
	PortDeposito = Conf.PortDeposito
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	Conf = config
	return config
}
