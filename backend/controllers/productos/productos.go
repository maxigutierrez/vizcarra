package productos

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
	Productos       []Models.Productos `json:"productos,omitempty"`
	Producto        *Models.Productos  `json:"producto,omitempty"`
	TotalDataSize int              `json:"totalDataSize,omitempty"`
}

func Create(c echo.Context) error {
	db := database.GetDb()
	productos := new(Models.Productos)
	if err := c.Bind(productos); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&productos).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	data := Data{Producto: productos}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Update(c echo.Context) error {
	db := database.GetDb()
	productos := new(Models.Productos)
	if err := c.Bind(productos); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body ",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := db.Save(&productos).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	data := Data{Producto: productos}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Data:    data,
		Message: "Los datos se actualizaron correctamente. ",
	})
}
func GetAll(c echo.Context) error {
	db := database.GetDb()

	db = db.Joins("JOIN marcas ON marcas.id = productos.marcas_id")
	if c.QueryParam("q") != "" {
		db = db.Where(` productos.id like ?	
		or producto like ?  
		or marcas.marca  like ? `,
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%",
			"%"+c.QueryParam("q")+"%")
	}

	//Order By
	if c.QueryParam("sortField") != "" {
		db = db.Order(c.QueryParam("sortField") + " " + c.QueryParam("sortOrder"))
	} else {
		db = db.Order("id")
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
	var productos []Models.Productos
	db.Offset(offset).Limit(limit).Preload("Marca").Find(&productos)
	db.Table("productos").Count(&totalDataSize)
	data := Data{Productos: productos, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	productos := new(Models.Productos)
	//db.Preload("Franquicias").Preload("Roles").Find(&productos, id)
	db.Find(&productos, id)

	data := Data{Producto: productos}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
func Delete(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec("DELETE FROM productos WHERE id = ?", c.Param("id")).Error; err != nil {
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
