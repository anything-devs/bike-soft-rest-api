package routes

import (
	"github.com/anything-devs/bike-soft-rest-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func Rutas(router *gin.Engine) {
	rutasInicio(router)
	rutasProductos(router)
}

func rutasInicio(router *gin.Engine) {
	router.GET("/", controllers.ControladorPaginaInicio)
}

/*
Metodo que contiene las rutas que se utilizan con productos
*/
func rutasProductos(router *gin.Engine) {
	router.GET("/productos", controllers.GetProductos)
}
