package controllers

import (
	"math"
	"net/http"
	"regexp"
	"strconv"

	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

const ERROR = "Error BD"

/*
* Método para obtener la lista de los productos
 */
func GetProductosAZ(ctx *gin.Context) {
	var productos []models.Producto
	if err := configs.BD.Order("nombre").Find(&productos).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ERROR: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, productos)
}

/*
* Método para obtener la lista de los productos
 */
func GetProductosZA(ctx *gin.Context) {
	var productos []models.Producto
	if err := configs.BD.Order("nombre DESC").Find(&productos).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ERROR: err.Error()})
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
* esto se debe hacer en la ruta http://localhost:8080/producto?nombre=&codigo=, puede ser cualquier parametro
 */
func GetProducto(ctx *gin.Context) {
	var productos []models.Producto
	var productoGet models.ProductoGet
	if ctx.ShouldBind(&productoGet) == nil {

		match, _ := regexp.MatchString("^[[:alpha:]]{3}[[:digit:]]{3}$", productoGet.Codigo)
		if productoGet.Codigo != "" && productoGet.Nombre == "" {
			if !match {
				ctx.JSON(http.StatusBadRequest, gin.H{"Error Codigo": "El código debe tener una longitud de 6 caracteres, 3 letras y 3 números"})
				return
			}
			if err := configs.BD.Where("codigo= ?", productoGet.Codigo).First(&productos).Error; err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{ERROR: "Producto no encontrado por código en BD"})
				return
			}
			ctx.JSON(http.StatusOK, productos)
			return
		} else {
			if err := configs.BD.Where("nombre LIKE ?", productoGet.Nombre+"%").Find(&productos).Error; err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{ERROR: "Producto no encontrado por nombre en BD"})
				return
			}
			ctx.JSON(http.StatusOK, productos)
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"Error faltan datos": "Debe escribir un código o nombre de producto"})
}

/*
* Método para actualizar la cantidad y/o el precio base de un producto en especifico
@param ctx: parametro del contexto del programa
@return el producto con la cantidad y/o el precio base actualizado
*/
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

	if actualizarProducto.PrecioBase < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El precio base no puede ser negativo"})
		return
	}

	if actualizarProducto.Cantidad != 0 {
		cantidadStr := strconv.Itoa(actualizarProducto.Cantidad)
		cantidad, err := strconv.Atoi(cantidadStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La cantidad debe ser numérica"})
			return
		}

		if cantidad < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La cantidad no puede ser negativa"})
			return
		}

		producto.Cantidad = int8(cantidad)
	}

	if actualizarProducto.PrecioBase != 0 {
		producto.Precio_base = actualizarProducto.PrecioBase
	}

	if err := configs.BD.Save(&producto).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, producto)
}

/*
Metodo que crea productos nuevos en la base de datos
@param ctx: parametro del contexto del programa
@return el nuevo producto creado
*/
func CrearProducto(ctx *gin.Context) {
	var producto models.NuevoProducto
	const IVA float64 = 1.19
	const G float64 = 0.75

	if err := ctx.ShouldBindJSON(&producto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error con los datos": err.Error()})
		return
	}

	match, _ := regexp.MatchString("^[[:alpha:]]{3}[[:digit:]]{3}$", producto.Codigo)
	if match {
		nuevoProducto := models.Producto{Codigo: producto.Codigo, Nombre: producto.Nombre,
			Precio_base: producto.Precio_base, Precio_venta: float32(math.Round((float64(producto.Precio_base) * IVA) / G)), Cantidad: producto.Cantidad, CategoriaID: producto.CategoriaID}

		if err := configs.BD.Where("codigo= ?", nuevoProducto.Codigo).First(&nuevoProducto).Error; err == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Producto ya existe en la BD"})
			return
		}

		if nuevoProducto.Precio_base < 0 || nuevoProducto.Cantidad < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "El precio base y la cantidad deben ser mayores a 0"})
			return
		}

		if err := configs.BD.Create(&nuevoProducto).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error al crear Producto": err.Error()})
			return
		} else {
			ctx.JSON(http.StatusOK, nuevoProducto)
		}

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error Codigo": "El código debe tener una longitud de 6 caracteres, 3 letras y 3 números"})
		return
	}

}

/*
Metodo que elimina productos de la base de datos
@param ctx: parametro del contexto del programa
@return producto con el id seleccionado eliminado
*/
func EliminarProducto(ctx *gin.Context) {
	// Obtener el ID del producto a eliminar desde los parámetros de la URL
	id := ctx.Param("id")

	var producto models.Producto
	if err := configs.BD.Where("id = ?", id).First(&producto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	if producto.Cantidad > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No se puede eliminar un producto con cantidad mayor a 0"})
		return
	}

	// Eliminar el producto de la base de datos
	if err := configs.BD.Delete(&producto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Enviar una respuesta HTTP en formato JSON indicando que el producto ha sido eliminado
	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado correctamente"})
}
