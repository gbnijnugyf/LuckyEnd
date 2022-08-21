package controller

import (
	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	user, err := model.GetUserInfo(UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询用户信息失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询用户信息成功", model.ViewUser{
		Email:     user.Email,
		Wechat:    user.Wechat,
		Telephone: user.Tel,
	}))
}
