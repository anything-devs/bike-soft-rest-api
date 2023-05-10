package routes

import (
	"github.com/anything-devs/bike-soft-rest-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func Rutas(router *gin.Engine) {
	rutasInicio(router)
	rutasProductos(router)
	rutasCategorias(router)
}

func rutasInicio(router *gin.Engine) {
	router.GET("/", controllers.ControladorPaginaInicio)
}

/*
* Método que contiene las rutas que se utilizan con productos
 */
func rutasProductos(router *gin.Engine) {
	router.GET("/productos-AZ", controllers.GetProductosAZ)
	router.GET("/productos-ZA", controllers.GetProductosZA)
	router.GET("/productos/filtradosBajasUnidades/:cantidad", controllers.FiltroBajasUnidades)
	router.GET("/producto", controllers.GetProducto)
	router.PUT("/productoActualizarStock/:id", controllers.ActualizarStock)
	router.POST("/productoCrear", controllers.CrearProducto)
	router.DELETE("/productoEliminar/:id", controllers.EliminarProducto)
}

/*
* Método que contiene las rutas que se utilizan con categorias
 */
func rutasCategorias(router *gin.Engine) {
	router.GET("/categorias", controllers.GetCategorias)
}
