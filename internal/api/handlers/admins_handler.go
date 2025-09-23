package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminsHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admins.html", nil)
}
