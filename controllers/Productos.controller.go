package controllers

import (
	"net/http"
	"strconv"

	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

/*
* Método para obtener la lista de los productos
 */
func GetProductos(ctx *gin.Context) {
	var productos []models.Producto
	if err := configs.BD.Find(&productos).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, productos)
}

/*
* Método para obtener la lista filtrada de productos con bajas unidades
 */
func FiltroBajasUnidades(ctx *gin.Context) {
	var productos []models.Producto
	var filtrado []models.Producto

	cantidad, err := strconv.Atoi(ctx.Param("cantidad"))
	if err != nil {
		// Manejo de error si el valor de cantidad no es un número válido
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "La cantidad debe ser un número entero válido"})
		return
	}

	// Manejo de la lista de productos
	if err := configs.BD.Find(&productos).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	/*
	* Manejo del filtrado por cantidad,
	* esto se debe hacer en la ruta http://localhost:8080/productos/filtradosBajasUnidades/
	* más el número de bajas unidades requeridos
	 */
	for _, producto := range productos {
		if int(producto.Cantidad) < cantidad {
			filtrado = append(filtrado, producto)
		}
	}

	// Manejo de la lista esta vacia, retorna la lista vacia
	if len(filtrado) == 0 {
		ctx.JSON(http.StatusOK, []models.Producto{})
		return
	}
	ctx.JSON(http.StatusOK, filtrado)
}

func ActualizarStock(ctx *gin.Context) {
	id := ctx.Param("id")

	// Obtener la nueva cantidad de la URL
	nuevaCantidad, err := strconv.Atoi(ctx.Param("cantidad"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "La cantidad debe ser un número entero válido"})
		return
	}

	// Buscar el producto en la base de datos y actualizar la cantidad
	var producto models.Producto
	if err := configs.BD.Where("id = ?", id).First(&producto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	producto.Cantidad = int8(nuevaCantidad)
	if err := configs.BD.Save(&producto).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el stock del producto"})
		return
	}
	ctx.JSON(http.StatusOK, producto)
}
