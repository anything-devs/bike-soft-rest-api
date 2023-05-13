package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anything-devs/bike-soft-rest-api.git/controllers"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/anything-devs/bike-soft-rest-api.git/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestCrearProductoCreado(t *testing.T) {
	mockRepository := &repositories.MockProductoRepository{}
	productoController := controllers.NewProductoController(mockRepository)

	requestBody := `{
		"codigo": "BIL0020",
		"nombre": "Prueba 20",
		"precio_base": 15000,
		"cantidad": 15,
		"categoria_id": 1
	}`

	request, _ := http.NewRequest(http.MethodPost, "/productoCrear", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = request

	productoController.CrearProducto(context)

	assert.Equal(t, http.StatusOK, context.Writer.Status())
}

func TestProductoEliminado(t *testing.T) {
	mockRepository := &repositories.MockProductoRepository{}
	productoController := controllers.NewProductoController(mockRepository)

	producto := &models.Producto{
		ID:           0,
		Codigo:       "BIL002",
		Nombre:       "Prueba 2",
		Precio_base:  15000,
		Precio_venta: 23800,
		Cantidad:     15,
		CategoriaID:  1,
	}
	mockRepository.CrearProducto(producto)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Params = append(context.Params, gin.Param{Key: "id", Value: "1"})
	productoController.EliminarProducto(context)

	assert.Equal(t, http.StatusOK, context.Writer.Status())
}
