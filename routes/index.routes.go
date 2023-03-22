package routes

import (
	"github.com/anything-devs/bike-soft-rest-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func Rutas(router *gin.Engine) {
	rutasInicio(router)
}

func rutasInicio(router *gin.Engine) {
	router.GET("/", controllers.ControladorPaginaInicio)
}
