package controller

import (
	"net/http"
	"test/common"
	"test/helper"
	"test/model"

	"github.com/gin-gonic/gin"
)

func CallBack(c *gin.Context) {

}

func CheckUserEmail(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	UserEmail := model.GetUserEmailByUserID(UserID)
	if UserEmail == "" {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "未绑定邮箱", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "邮箱已绑定", UserEmail))
}

func BindEmail(c *gin.Context) {
	var user model.User
	student_number := c.MustGet("student_number").(string)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
		return
	}
	user.IdcardNumber = student_number
	if res := model.BindEmail(user); res.Status == common.CodeError {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "邮箱绑定失败", res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "绑定邮箱成功", user.Email))
}
