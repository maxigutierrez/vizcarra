package clientes

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
	Clientes       []Models.Clientes `json:"clientes,omitempty"`
	Cliente        *Models.Clientes  `json:"cliente,omitempty"`
	TotalDataSize int              `json:"totalDataSize,omitempty"`
}

func Create(c echo.Context) error {
	db := database.GetDb()
	clientes := new(Models.Clientes)
	if err := c.Bind(clientes); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&clientes).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	data := Data{Cliente: clientes}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Update(c echo.Context) error {
	db := database.GetDb()
	clientes := new(Models.Clientes)
	if err := c.Bind(clientes); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body ",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := db.Save(&clientes).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	data := Data{Cliente: clientes}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Data:    data,
		Message: "Los datos se actualizaron correctamente. ",
	})
}
func GetAll(c echo.Context) error {
	db := database.GetDb()

	// db = db.Joins("JOIN marcas ON marcas.id = clientes.marcas_id")
	if c.QueryParam("q") != "" {
		db = db.Where(` clientes.id like ?	
		or cliente like ?  
		or dni like ? 
		or celular like ? 
		or domicilio like ? 
		or localidad like ? `,
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%")
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
	var clientes []Models.Clientes
	db.Offset(offset).Limit(limit).Find(&clientes)
	db.Table("clientes").Count(&totalDataSize)
	data := Data{Clientes: clientes, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	clientes := new(Models.Clientes)
	//db.Preload("Franquicias").Preload("Roles").Find(&clientes, id)
	db.Find(&clientes, id)

	data := Data{Cliente: clientes}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Delete(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec("DELETE FROM clientes WHERE id = ?", c.Param("id")).Error; err != nil {
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
