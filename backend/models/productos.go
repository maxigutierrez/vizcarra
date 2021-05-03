package models

type Productos struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Producto  string `json:"producto"`
}

func (Productos) TableName() string {
	return "productos"
}
