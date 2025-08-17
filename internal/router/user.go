package router

import (
	"gin_boot/internal/controller"
	"gin_boot/internal/dao"
	"gin_boot/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRouter(router *gin.RouterGroup, db *gorm.DB) {
	udao := dao.NewUserDao(db)
	usvc := service.NewUserService(udao)
	uc := controller.NewUserController(usvc)
	user := router.Group("/users")
	{
		user.GET("/login", uc.Login)
	}

}
