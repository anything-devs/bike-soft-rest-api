package controllers

import (
	"net/http"

	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

/*
* Método para obtener la lista de los categorias
 */
func GetCategorias(ctx *gin.Context) {
	var categorias []models.Categoria
	if err := configs.BD.Find(&categorias).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ERROR: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categorias)
}
