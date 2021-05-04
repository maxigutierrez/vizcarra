package models

type Productos struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Producto  string `json:"producto"`
	Marcas_id uint   `json:"marcas_id"`
	Marca     Marcas `json:"marca" gorm:"ForeignKey:marcas_id;AssociationForeignKey:id"`
}

func (Productos) TableName() string {
	return "productos"
}
