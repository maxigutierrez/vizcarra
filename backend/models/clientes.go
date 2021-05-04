package models

type Clientes struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Cliente   string `json:"cliente"`
	Dni 	  uint   `json:"dni"`
	Celular	  string `json:"celular"`
	Domicilio string `json:"domicilio"`
	Localidad string `json:"localidad"`
	// Marca     Marcas `json:"marca" gorm:"ForeignKey:marcas_id;AssociationForeignKey:id"`
}

func (Clientes) TableName() string {
	return "clientes"
}
