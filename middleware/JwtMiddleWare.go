package middleware

import (
	"log"
	"net/http"
	"test/common"
	"test/controller"
	"test/helper"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeError, "token不存在", nil))
		c.Abort()
		return
	}
	student_number, err := helper.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpries, "权限不足", nil))
		c.Abort()
		return
	}
	UserID := controller.GeuUserIDFromDB(student_number)
	if UserID == common.CodeError {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpries, "权限不足", nil))
		log.Println("=========异常登录记录==========")
		log.Println(student_number)
		c.Abort()
		return
	}
	log.Print(UserID)
	c.Set("student_number", student_number)
	c.Set("user_id", UserID)
	c.Next()
}
