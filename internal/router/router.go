package router

import "github.com/gin-gonic/gin"

func InitRouter(server *gin.Engine) {
	baseRouter := server.Group("/api/v1")
	InitUserRouter(baseRouter)
}
