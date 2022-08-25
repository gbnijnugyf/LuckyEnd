package middleware

import (
	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"
	"log"
	"net/http"

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
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpires, "权限不足", nil))
		c.Abort()
		return
	}
	UserID, _ := model.GetUserIDByStudentNumber(student_number)
	if UserID == common.CodeError {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpires, "权限不足", nil))
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
