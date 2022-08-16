package router

import (
	"test/controller"
	"test/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.POST("api/whutlogin", controller.Login)

	api := r.Group("api/")
	api.Use(middleware.AuthMiddleware)
	{
		api.GET("/user/info/wishman", controller.GetUserInfo)
		api.GET("/user/info/lightman", controller.GetUserInfo)

		wishes := api.Group("/wishes")
		{
			wishes.POST("/add", controller.UserAddDesire)
			wishes.POST("/light", controller.UserLightDesire)
			wishes.POST("/achieve", controller.AchieveUserDesire)
			wishes.GET("/user/post", controller.GetUserCreateDesires)
			wishes.GET("/user/light", controller.GetUserLightDesires)
			wishes.GET("/categories", controller.GetUserDesireByType)
			wishes.DELETE("", controller.DeleteUserDesire)
			wishes.POST("/giveup", controller.CancelUserLight)
		}

		api.GET("")
	}
}
