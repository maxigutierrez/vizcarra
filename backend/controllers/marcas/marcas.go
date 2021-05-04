package marcas

import (
	"net/http"

	"../../database"
	Models "../../models"
	"../../utils"
	"github.com/labstack/echo"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    Data   `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type Data struct {
	Marcas       []Models.Marcas `json:"marcas,omitempty"`
	Marca        *Models.Marcas  `json:"marca,omitempty"`
	TotalDataSize int            `json:"totalDataSize,omitempty"`
}

func Create(c echo.Context) error {
	db := database.GetDb()
	marcas := new(Models.Marcas)
	if err := c.Bind(marcas); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&marcas).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	data := Data{Marca: marcas}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Update(c echo.Context) error {
	db := database.GetDb()
	marcas := new(Models.Marcas)
	if err := c.Bind(marcas); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body ",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := db.Omit("alta", "ualta", "baja", "ubaja").Save(&marcas).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	data := Data{Marca: marcas}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Data:    data,
		Message: "Los datos se actualizaron correctamente. ",
	})
}
func GetAll(c echo.Context) error {
	db := database.GetDb()

	if c.QueryParam("q") != "" {
		db = db.Where("id like ? ", "%"+c.QueryParam("q")+"%").
			Or("marca like ? ", "%"+c.QueryParam("q"))
	}

	//Order By
	if c.QueryParam("sortField") != "" {
		db = db.Order(c.QueryParam("sortField") + " " + c.QueryParam("sortOrder"))
	} else {
		db = db.Order("id desc")
	}

	//Paginacion
	//===============================================
	var page uint = 1
	var limit uint = 10
	var offset uint = 0
	var totalDataSize int = 0
	if c.QueryParam("limit") != "" {
		limit = utils.ParseInt(c.QueryParam("limit"))
	}
	if c.QueryParam("page") != "" {
		page = utils.ParseInt(c.QueryParam("page"))
	}
	offset = limit * (page - 1)
	//===============================================
	var marcas []Models.Marcas
	db.Offset(offset).Limit(limit).Find(&marcas)
	db.Table("marcas").Count(&totalDataSize)
	data := Data{Marcas: marcas, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	marcas := new(Models.Marcas)
	//db.Preload("Franquicias").Preload("Roles").Find(&marcas, id)
	db.Find(&marcas, id)

	data := Data{Marca: marcas}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Delete(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec("DELETE FROM marcas WHERE id = ?", c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Message: "Ok",
	})
}
