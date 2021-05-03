package google

import (
	"net/http"
	"github.com/labstack/echo"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func Get(c echo.Context) error {

	url := c.Param("url")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return 1
}

