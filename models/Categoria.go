package models

import "gorm.io/gorm"

/*
* Modelo de base de datos para la tabla categorias
 */
type Categoria struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"type:varchar(100);not null" json:"nombre"`
	Descripcion string `gorm:"type:varchar(200);not null" json:"descripcion"`
	Productos   []Producto
}
