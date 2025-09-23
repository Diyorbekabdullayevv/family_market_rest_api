package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CatalogHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "catalog.html", nil)
}

func CartHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cart.html", nil)
}
