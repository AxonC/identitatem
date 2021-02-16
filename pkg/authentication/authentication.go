package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewIdentifyProvider() *gin.Engine {
	router := gin.Default()

	router.GET("/health_check", healthCheckHandler)

	return router
}

func healthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
