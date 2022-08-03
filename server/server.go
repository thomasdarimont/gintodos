package server

import (
	"gintodos/api"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	api.RegisterTodoEndpoints(r)
	api.RegisterMonitoringEndpoints(r)
	_ = r.Run("localhost:9090")
}
