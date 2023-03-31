package controllers

import (
	"net/http"

	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

/*
Metodo para obtener la lista de los productos
*/
func GetProductos(ctx *gin.Context) {
	var productos []models.Producto
	if err := configs.BD.Find(&productos).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, productos)
}
