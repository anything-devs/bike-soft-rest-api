package models

import "gorm.io/gorm"

/*
* Modelo de base de datos para la tabla productos
 */
type Producto struct {
	gorm.Model

	ID           uint    `gorm:"primaryKey"`
	Codigo       string  `gorm:"type:varchar(10);not null;unique_index" json:"codigo" binding:"required"`
	Nombre       string  `gorm:"type:varchar(100);not null" json:"nombre" binding:"required"`
	Precio_base  float32 `gorm:"type:float;not null" json:"precio_base" binding:"required"`
	Precio_venta float32 `gorm:"type:float;not null" json:"precio_venta" binding:"required"`
	Cantidad     int8    `gorm:"type:int;" json:"cantidad" binding:"required"`
}

/*
Struct para realizar las consultas por nombre o codigo de un producto especifico
*/
type ProductoGet struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}

type ActualizarProducto struct {
	Cantidad int `json:"cantidad"`
}
