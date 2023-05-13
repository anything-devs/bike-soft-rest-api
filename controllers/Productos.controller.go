package controllers

import (
	"math"
	"net/http"
	"regexp"
	"strconv"

	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/anything-devs/bike-soft-rest-api.git/repositories"
	"github.com/gin-gonic/gin"
)

const ERROR = "Error BD"

// ProductoController es el controlador de productos
type ProductoControllerImpl struct {
	productoRepository repositories.ProductoRepository
}

// ProductoController es la interfaz de ProductoController
type ProductoController interface {
	CrearProducto(ctx *gin.Context)
	EliminarProducto(ctx *gin.Context)
}

/*
* Constructor de ProductoController
* @param productoRepository: repositorio de productos
* @return instancia de ProductoController
 */
func NewProductoController(productoRepository repositories.ProductoRepository) ProductoController {
	return &ProductoControllerImpl{productoRepository: productoRepository}
}

/*
* Método para obtener la lista de los productos
 */
func GetProductosAZ(ctx *gin.Context) {
	var productos []models.Producto
	if err := configs.ConectarBD().Order("nombre").Find(&productos).Error; err != nil {
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
	if err := configs.ConectarBD().Order("nombre DESC").Find(&productos).Error; err != nil {
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
	if err := configs.ConectarBD().Find(&productos).Error; err != nil {
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
			if err := configs.ConectarBD().Where("codigo= ?", productoGet.Codigo).First(&productos).Error; err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{ERROR: "Producto no encontrado por código en BD"})
				return
			}
			ctx.JSON(http.StatusOK, productos)
			return
		} else {
			if err := configs.ConectarBD().Where("nombre LIKE ?", productoGet.Nombre+"%").Find(&productos).Error; err != nil {
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
	if err := configs.ConectarBD().Where("id = ?", id).First(&producto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	var actualizarProducto models.ActualizarProducto

	if err := ctx.ShouldBindJSON(&actualizarProducto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if actualizarProducto.Cantidad >= 0 {
		producto.Cantidad = int8(actualizarProducto.Cantidad)
	}

	if actualizarProducto.PrecioBase > 0 {
		producto.Precio_base = actualizarProducto.PrecioBase
	}

	if actualizarProducto.PrecioBase < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El precio base no puede ser negativo"})
		return
	}

	if actualizarProducto.Cantidad < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "La cantidad no puede ser negativa"})
		return
	}

	if err := configs.ConectarBD().Save(&producto).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, producto)
}

/*
* Método para crear un producto
@param ctx: parametro del contexto del programa
*/
func (cp *ProductoControllerImpl) CrearProducto(ctx *gin.Context) {
	var producto models.NuevoProducto
	const IVA float64 = 1.19
	const GANANCIA float64 = 0.75

	if err := ctx.ShouldBindJSON(&producto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error con los datos": err.Error()})
		return
	}

	match, _ := regexp.MatchString("^[[:alpha:]]{3}[[:digit:]]{3}$", producto.Codigo)
	if match {
		nuevoProducto := models.Producto{
			Codigo:       producto.Codigo,
			Nombre:       producto.Nombre,
			Precio_base:  producto.Precio_base,
			Precio_venta: float32(math.Round((float64(producto.Precio_base) * IVA) / GANANCIA)),
			Cantidad:     producto.Cantidad,
			CategoriaID:  producto.CategoriaID,
		}

		if nuevoProducto.Precio_base < 0 || nuevoProducto.Cantidad < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "El precio base y la cantidad deben ser mayores a 0"})
			return
		}

		if err := cp.productoRepository.CrearProducto(&nuevoProducto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{ERROR: err.Error()})
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
* Método para eliminar un producto en especifico
@param ctx: parametro del contexto del programa
*/
func (cp *ProductoControllerImpl) EliminarProducto(ctx *gin.Context) {
	id := ctx.Param("id")
	var producto models.Producto

	if producto.Cantidad > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No se puede eliminar un producto con cantidad mayor a 0"})
		return
	}

	if err := cp.productoRepository.EliminarProducto(id, &producto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado correctamente"})
}
