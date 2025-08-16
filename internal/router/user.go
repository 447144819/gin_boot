package router

import "github.com/gin-gonic/gin"

func InitUserRouter(router *gin.RouterGroup) {
	//userHandler := web.NewUserHandler()
	user := router.Group("/users")
	{
		user.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  "success",
			})
		})
	}

}
