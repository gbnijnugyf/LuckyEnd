package controller

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"
	"github.com/shawu21/test/service"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	userLogin := &model.UserLogin{}
	if err := c.ShouldBindJSON(userLogin); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
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
	//TODO data may be delay from auth when user update his data in auth
	user, err := model.UserCheck(userLogin.Email)
	if err != nil {
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
	token, err := helper.CreatToken(user.IdcardNumber)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建token失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "登录成功", token))
}
