package router

import (
	"net/http"
	"strings"

	"github.com/shawu21/LuckyBackend/controller"
	"github.com/shawu21/LuckyBackend/helper"
	"github.com/shawu21/LuckyBackend/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.POST("api/whutlogin", controller.Login)
	r.GET("api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("api/")
	api.Use(middleware.AuthMiddleware)
	{
		api.GET("/user/info", controller.GetUserInfo)

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

	}

	// jwt test
	r.GET("test1", func(c *gin.Context) {
		token, _ := helper.CreatToken("123456789")
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	r.GET("test2", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		prefix := "Bearer"
		if token == "" || !strings.HasPrefix(token, prefix) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token不存在",
			})
			return
		}
		token = token[len(prefix)+1:]
		studentNumber, _ := helper.VerifyToken(token)
		c.JSON(http.StatusOK, gin.H{
			"token": studentNumber,
		})
	})
}
