package models

type Marcas struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Marca	  string `json:"marca"`
}

func (Marcas) TableName() string {
	return "marcas"
}
