package controller

import (
	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"
	"github.com/shawu21/test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLogin *model.UserLogin
	var user *model.User
	if err := c.ShouldBindJSON(userLogin); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", nil))
		return
	}
	accessToken, err := service.SendForm(userLogin.Email, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "登录失败", err))
		return
	}
	if accessToken == "" {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "用户名不存在或密码错误", nil))
		return
	}
	if err := model.UserCheck(userLogin.Email); err != nil {
		user, err = service.GetInfo(accessToken)
		if err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "获取用户信息失败", nil))
			return
		}
		err = model.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建用户失败", nil))
			return
		}
	}
	myToken, err := helper.CreatToken(user.IdcardNumber)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建token失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "创建用户成功", myToken))
}

// func CheckUserEmail(c *gin.Context) {
// 	UserID := c.MustGet("user_id").(int)
// 	UserEmail := model.GetUserEmailByUserID(UserID)
// 	if UserEmail == "" {
// 		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "未绑定邮箱", nil))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "邮箱已绑定", UserEmail))
// }

// func BindEmail(c *gin.Context) {
// 	var user model.User
// 	student_number := c.MustGet("student_number").(string)
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
// 		return
// 	}
// 	user.IdcardNumber = student_number
// 	if res := model.BindEmail(user); res.Status == common.CodeError {
// 		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "邮箱绑定失败", res.Data))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "绑定邮箱成功", user.Email))
// }
