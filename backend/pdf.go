package main

import (
	"./config"
	"./database"
  "os"
  ComprobantesTiposController "./controllers/comprobantes_tipos"
)

func main() {
	config.Init()
	database.InitDb()
  id:=	os.Args[1]
	path:=	os.Args[2]
	ComprobantesTiposController.PDF(id,path)
}
