package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	ProductsRouter(server)
}
