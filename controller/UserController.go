package controller

import (
	"net/http"
	"strconv"

	"github.com/shawu21/LuckyBackend/common"
	"github.com/shawu21/LuckyBackend/helper"
	"github.com/shawu21/LuckyBackend/model"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	UserID := c.Query("id")
	userID, err := strconv.Atoi(UserID)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "id is not right", nil))
		return
	}
	if userID == common.CheckSelf {
		id := c.MustGet("user_id").(int)
		user, err := model.GetUserInfo(id)
		if err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "查询用户信息失败", nil))
			return
		}
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询用户信息成功", model.ViewUser{
			Name:   user.Name,
			QQ:     user.QQ,
			Email:  user.Email,
			Wechat: user.Wechat,
			Tel:    user.Tel,
		}))
		return
	}
	user, err := model.GetUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "查询用户信息失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询用户信息成功", model.ViewUser{
		Name:   user.Name,
		QQ:     user.QQ,
		Email:  user.Email,
		Wechat: user.Wechat,
		Tel:    user.Tel,
	}))
}
