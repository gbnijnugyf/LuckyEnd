package router

import (
	"github.com/shawu21/test/controller"
	"github.com/shawu21/test/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.POST("api/whutlogin", controller.Login)

	api := r.Group("api/")
	api.Use(middleware.AuthMiddleware)
	{
		api.GET("/user/info/wishman", controller.GetUserInfo)
		api.GET("/user/info/lightman", controller.GetUserInfo)

		wishes := api.Group("/desires")
		{
			wishes.POST("/add", controller.UserAddDesire)
			wishes.POST("/light", controller.UserLightDesire)
			wishes.POST("/achieve", controller.AchieveUserDesire)
			wishes.GET("/details", controller.DesireDetail)
			wishes.GET("/user/post", controller.GetUserCreateDesires)
			wishes.GET("/user/light", controller.GetUserLightDesires)
			wishes.GET("/categories", controller.GetUserDesireByType)
			wishes.DELETE("/delete", controller.DeleteUserDesire)
			wishes.POST("/giveup", controller.CancelUserLight)
		}

		api.GET("")
	}
}
