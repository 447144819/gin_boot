package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(server *gin.Engine, db *gorm.DB) {
	baseRouter := server.Group("/api/v1")
	InitUserRouter(baseRouter, db)
}
