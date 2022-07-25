package server

import (
	"gintodos/api"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	api.RegisterTodoAPI(router)

	router.Run("localhost:9090")
}
