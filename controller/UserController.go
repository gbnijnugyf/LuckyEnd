package controller

import (
	"net/http"
	"test/common"
	"test/helper"
	"test/model"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	user, err := model.GetUserInfo(UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询用户信息失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询用户信息成功", *user))
}


