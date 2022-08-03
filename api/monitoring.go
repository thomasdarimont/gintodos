package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterMonitoringEndpoints(router *gin.Engine) {
	router.GET("/metrics", getMetrics)
	router.GET("/health", getHealth)
}

//
// curl -v http://localhost:9090/todos
func getMetrics(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, map[string]any{
		"metrics": "OK",
	})
}

func getHealth(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, map[string]any{
		"health": "OK",
	})
}
