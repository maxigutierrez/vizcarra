package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../database"
	Models "../models"
	"./../config"
)

type Guias struct {
	Data []Guia `json:"data"`
}
type Guia struct {
	Guia_nro  string `json:"guia_nro"`
	Bulto_nro string `json:"bulto_nro"`
	Estado    string `json:"estado"`
}
type Bultos struct {
	Data []Bulto `json:"data"`
}
type Bulto struct {
	Tagid     string `json:"guia_nro"`
	Bulto_nro string `json:"bulto_nro"`
}
type Resultado struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Debug   string `json:"debug"`
	Code    int    `json:"code"`
}

func main() {
	fmt.Println("inicial...")
	config.Init()
	database.InitDb()
	var bultos []Models.Comprobantes_detalles_bultos
	db := database.GetDb()
	db.Raw("Select top 1000 * from Comprobantes_detalles_bultos where sincro is null and len(etiqueta)>5 ").Preload("Detalle").Find(&bultos)

	for index := 0; index < len(bultos); index++ {
		var guias []Guia
		guia := new(Guia)
		guia.Guia_nro = fmt.Sprint(bultos[index].Detalle.Comprobantes_id)
		guia.Bulto_nro = bultos[index].Etiqueta
		guia.Estado = "1"
		guias = append(guias, *guia)
		fmt.Println("Send :", bultos[index].Etiqueta, " ", fmt.Sprint(bultos[index].Detalle.Comprobantes_id))
		g := new(Guias)
		g.Data = guias
		if Bluit_guia(*g) {
			if err := db.Exec("update comprobantes_detalles_bultos set sincro = 1 where id = ?", bultos[index].ID).Error; err != nil {
				fmt.Println("Error :", bultos[index].Etiqueta)
			}
		}
	}
}

func Bluit_guia(guias Guias) bool {
	//url := "https://credifin.blue-it.com.ar/ws/guias"
	  url :="http://credifintest.blue-it.com.ar/ws/guias"
	method := "POST"
	client := &http.Client{Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(guias)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDA1NTgsImlkcGFzcyI6IiUhcyh1aW50PTkxKSIsImlkdXN1YXJpbyI6IjYifQ.T5VFsDpGcxrIC7xv3A9WfovxT9DnvOMDwGBclmc61cU")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	resultado := new(Resultado)

	if err := json.Unmarshal(body, &resultado); err != nil {
		fmt.Sprintln("Error -%+v", err)
		return false
	}
	if !resultado.Success {
		fmt.Print("Resultado: ", resultado.Success, "Error: ", resultado.Msg)
		if resultado.Msg == "Bulto Existente" {
			return true
		}
	} else {
		fmt.Print("Resultado: ", resultado.Success)
	}
	return resultado.Success
}
func Bluit_bultos(guias Bultos) {

	//url := "https://credifin.blue-it.com.ar/ws/bultos"
	  url :="http://credifintest.blue-it.com.ar/ws/guias"
	method := "POST"
	client := &http.Client{Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(guias)
	fmt.Printf("%+v", guias)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDA1NTgsImlkcGFzcyI6IiUhcyh1aW50PTkxKSIsImlkdXN1YXJpbyI6IjYifQ.T5VFsDpGcxrIC7xv3A9WfovxT9DnvOMDwGBclmc61cU")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
