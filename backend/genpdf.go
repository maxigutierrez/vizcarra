package main

import (
	"os"

	"./config"
	ComprobantesTiposController "./controllers/comprobantes_tipos"
	"./database"
)

func main() {
	config.Init()
	database.InitDb()
	cuit := os.Args[1]
	fechadesde := os.Args[2]
	fechahasta := os.Args[3]
	ComprobantesTiposController.GenPdf(cuit, fechadesde, fechahasta)
}
