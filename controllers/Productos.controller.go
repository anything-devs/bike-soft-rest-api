package controllers

import (
	"net/http"
	"regexp"
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
	if err := configs.BD.Order("nombre").Find(&productos).Error; err != nil {
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

/*
* Método para obtener un producto por código o nombre especifico
* esto se debe hacer en la ruta http://localhost:8080/producto
 */
func GetProducto(ctx *gin.Context) {
	var producto models.Producto
	var productoGet models.ProductoGet
	if ctx.BindJSON(&productoGet) == nil {
		//Expresión regular para validación de caracteristicas del codigo del producto
		match, _ := regexp.MatchString("^[[:alpha:]]{3}[[:digit:]]{3}$", productoGet.Codigo)
		if productoGet.Codigo != "" && productoGet.Nombre == "" {
			if !match {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "El código debe tener una longitud de 6 caracteres, 3 letras y 3 números"})
				return
			}
			if err := configs.BD.Where("codigo= ?", productoGet.Codigo).First(&producto).Error; err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado por código"})
				return
			}
			ctx.JSON(http.StatusOK, producto)
			return
		} else {
			if err := configs.BD.Where("nombre= ?", productoGet.Nombre).First(&producto).Error; err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado por nombre"})
				return
			}
			ctx.JSON(http.StatusOK, producto)
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Debe escribir un código o nombre de producto"})
}

func ActualizarStock(ctx *gin.Context) {
	id := ctx.Param("id")

	// Buscar el producto en la base de datos y actualizar la cantidad
	var producto models.Producto
	if err := configs.BD.Where("id = ?", id).First(&producto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	var actualizarProducto models.ActualizarProducto

	if err := ctx.ShouldBindJSON(&actualizarProducto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if actualizarProducto.PrecioBase != 0 {
		producto.Precio_base = actualizarProducto.PrecioBase
	}

	if err := configs.BD.Model(&producto).Updates(models.Producto{Cantidad: int8(actualizarProducto.Cantidad)}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if actualizarProducto.Cantidad > 0 {
		producto.Cantidad = int8(actualizarProducto.Cantidad)
	}

	if err := configs.BD.Model(&producto).Updates(models.Producto{Precio_base: float32(actualizarProducto.PrecioBase)}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, producto)
}
