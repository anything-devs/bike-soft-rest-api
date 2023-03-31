package models

import "gorm.io/gorm"

/*
Modelo de base de datos para la tabla productos
*/
type Producto struct {
	gorm.Model

	ID           uint    `gorm:"primaryKey"`
	Codigo       string  `gorm:"type:varchar(10);not null;unique_index" json:"codigo"`
	Nombre       string  `gorm:"type:varchar(100);not null" json:"nombre"`
	Precio_base  float32 `gorm:"type:float;not null" json:"precio_base"`
	Precio_venta float32 `gorm:"type:float;not null" json:"precio_venta"`
	Cantidad     int8    `gorm:"type:int;" json:"cantidad"`
}
