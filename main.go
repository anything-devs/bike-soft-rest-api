package main

import (
	"github.com/anything-devs/bike-soft-rest-api.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Rutas(router)

	router.Run(":8080")
}
