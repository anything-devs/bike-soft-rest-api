package controllers

import "github.com/gin-gonic/gin"

func ControladorPaginaInicio(ctx *gin.Context) {
	ctx.String(200, "Hola Anything Devs")
}
