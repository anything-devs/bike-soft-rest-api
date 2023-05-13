package controllers

import (
	"net/http"

	"github.com/anything-devs/bike-soft-rest-api.git/repositories"
	"github.com/gin-gonic/gin"
)

// CategoriaController es el controlador de categorias
type CategoriaControllerImpl struct {
	categoriaRepository repositories.CategoriaRepository
}

// CategoriaController es la interfaz de CategoriaController
type CategoriaController interface {
	GetCategorias(ctx *gin.Context)
}

func NewCategoriaController(categoriaRepository repositories.CategoriaRepository) CategoriaController {
	return &CategoriaControllerImpl{categoriaRepository: categoriaRepository}
}

func (cc *CategoriaControllerImpl) GetCategorias(ctx *gin.Context) {
	categorias, err := cc.categoriaRepository.GetCategorias()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categorias)
}
