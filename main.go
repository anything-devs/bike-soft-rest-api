package main

import (
	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/anything-devs/bike-soft-rest-api.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConectarBD()
	configs.BD.AutoMigrate(models.Producto{})

	router := gin.Default()

	routes.Rutas(router)

	router.Run(":8080")
}
