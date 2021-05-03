
package main

import (
	"net/http"
	UsuariosController "./controllers/usuarios"

	"fmt"
  "time"
	"strings"
  	"io/ioutil"

	l "./logs"
	"./database"
  "./config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
type ResponseMessage struct {
	Status  string `json:"status"`

	Message string `json:"message,omitempty"`
}
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
	b.POST("/test", Test)



	e.Logger.Fatal(e.Start(config.PortSincro))

}
func Test(c echo.Context) error {
		for i:=0;i<100000000000000000;i++{

		}
		bluit_guia(1234, "1234", "")
		return c.JSON(http.StatusOK, ResponseMessage{
			Status: "success",
		})
}
func bluit_guia(guia uint, etiqueta string, tag string) {

	sguia := fmt.Sprint(guia)
	url := "https://credifin.blue-it.com.ar/ws/guias"
	method := "POST"

	payload := strings.NewReader("{\n    \"data\": [\n        {\n            \"guia_nro\": \"" + sguia + "\",\n            \"bulto_nro\": \"" + etiqueta + "\"\n        }\n    ]\n}")

	client := &http.Client{Timeout: 1 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDA1NTgsImlkcGFzcyI6IiUhcyh1aW50PTkxKSIsImlkdXN1YXJpbyI6IjYifQ.T5VFsDpGcxrIC7xv3A9WfovxT9DnvOMDwGBclmc61cU")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	fmt.Println("Fin Guia")
}
