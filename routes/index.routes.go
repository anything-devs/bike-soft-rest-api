package routes

import (
	"github.com/anything-devs/bike-soft-rest-api.git/controllers"
	"github.com/anything-devs/bike-soft-rest-api.git/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
* Método que contiene las rutas de la aplicación
* @param router: enrutador de gin
* @param DB: conexión a la base de datos
 */
func Rutas(router *gin.Engine, DB *gorm.DB) {
	productoRepository := repositories.NewProductoRepository(DB)
	productoController := controllers.NewProductoController(productoRepository)
	categoriaRepository := repositories.NewCategoriaRepository(DB)
	categoriaController := controllers.NewCategoriaController(categoriaRepository)

	rutasInicio(router)
	rutasProductos(productoController, router)
	rutasCategorias(categoriaController, router)
}

func rutasInicio(router *gin.Engine) {
	router.GET("/", controllers.ControladorPaginaInicio)
}

/*
* Método que contiene las rutas que se utilizan con productos
* @param productoController: controlador de productos
 */
func rutasProductos(productoController controllers.ProductoController, router *gin.Engine) {
	router.GET("/productos-AZ", controllers.GetProductosAZ)
	router.GET("/productos-ZA", controllers.GetProductosZA)
	router.GET("/productos/filtradosBajasUnidades/:cantidad", controllers.FiltroBajasUnidades)
	router.GET("/producto", controllers.GetProducto)
	router.PUT("/productoActualizarStock/:id", controllers.ActualizarStock)
	router.POST("/productoCrear", productoController.CrearProducto)
	router.DELETE("/productoEliminar/:id", productoController.EliminarProducto)
}

/*
* Método que contiene las rutas que se utilizan con categorias
* @param categoriaController: controlador de categorias
 */
func rutasCategorias(categoriaController controllers.CategoriaController, router *gin.Engine) {
	router.GET("/categorias", categoriaController.GetCategorias)
}
