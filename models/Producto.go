package models

import "gorm.io/gorm"

/*
* Modelo de base de datos para la tabla productos
 */
type Producto struct {
	gorm.Model

	ID           uint    `gorm:"primaryKey"`
	Codigo       string  `gorm:"type:varchar(10);not null;unique_index" json:"codigo"`
	Nombre       string  `gorm:"type:varchar(100);not null" json:"nombre"`
	Precio_base  float32 `gorm:"type:float;not null" json:"precio_base"`
	Precio_venta float32 `gorm:"type:float;not null" json:"precio_venta"`
	Cantidad     int8    `gorm:"type:int;" json:"cantidad"`
	CategoriaID  uint    `gorm:"type:int;not null" json:"categoria_id"`
}

/*
Struct para realizar las consultas por nombre o codigo de un producto especifico
*/
type ProductoGet struct {
	Codigo string `form:"codigo"`
	Nombre string `form:"nombre"`
}

type ActualizarProducto struct {
	Cantidad   int     `json:"cantidad"`
	PrecioBase float32 `json:"precio_base"`
}

type NuevoProducto struct {
	Codigo      string  `json:"codigo" binding:"required"`
	Nombre      string  `json:"nombre" binding:"required"`
	Precio_base float32 `json:"precio_base" binding:"required"`
	Cantidad    int8    `json:"cantidad" binding:"required"`
	CategoriaID uint    `json:"categoria_id" binding:"required"`
}
