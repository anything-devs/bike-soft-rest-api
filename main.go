package main

import (
	"github.com/anything-devs/bike-soft-rest-api.git/configs"
	"github.com/anything-devs/bike-soft-rest-api.git/models"
	"github.com/anything-devs/bike-soft-rest-api.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	DB := configs.ConectarBD()
	DB.Table("productos").AutoMigrate(models.Producto{})
	DB.Table("categorias").AutoMigrate(models.Categoria{})

	routes.Rutas(router, DB)

	router.Run(":8080")
}
