package models

type Productos struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Producto  string `json:"producto"`
	Marcas_id uint   `json:"marcas_id"`
}

func (Productos) TableName() string {
	return "productos"
}
