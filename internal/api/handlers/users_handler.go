package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAccount(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_account.html", nil)
}
