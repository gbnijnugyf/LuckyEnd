package controller

import (
	"net/http"
	"test/common"
	"test/helper"
	"test/model"
	"time"

	"github.com/gin-gonic/gin"
)

func UserAddDesire(c *gin.Context) {
	var desireJson model.Desire

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	// desireMap := helper.Struct2Map(desireJson)

	// if res, err := desireValidate.ValidateMap(desireMap, "add"); !res {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
	// 	return
	// }
	UserID := c.MustGet("user_id").(int)
	School := c.MustGet("school").(int)
	desireJson.UserID = UserID
	desireJson.School = School

	// 检查用户当前许愿次数
	WishCount := model.GetUserWishCount(UserID)
	// 判断许愿总的次数是否超过上限
	if WishCount >= common.MaxWishCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "许愿次数已达上限", nil))
		return
	}

	desireJson.CreatAt = time.Now().In(common.ChinaTime)
	res := model.AddDesire(desireJson)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 用户点亮他人愿望
func UserLightDesire(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	DesireID := c.MustGet("wish_id").(int)
	LightCount := model.GetUserLightCount(UserID)
	// 判断点亮次数是否达到上限
	if LightCount >= common.MaxLightCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "点亮愿望次数已达上限", nil))
		return
	}
	LightCount = model.GetUserLightAtSameTimeCount(UserID)
	// 判断同时点亮次数是否达到上限
	if LightCount >= common.MaxLightSameCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "同时点亮愿望次数已达上限", nil))
		return
	}
	res := model.LightDesire(DesireID, UserID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 获取用户所有的愿望
func GetUserAllDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res := model.GetUserAllDesire(UserID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 获取用户投递的愿望
func GetUserCreateDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res := model.GetUserCreateDesire(UserID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 获取用户点亮的愿望
func GetUserLightDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res := model.GetUserLightDesire(UserID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 通过点击分类查看愿望
func GetUserDesireByTypeZ(c *gin.Context) {
	Type := c.MustGet("type").(int)
	res := model.GetWishByCategories(Type)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 用户删除愿望
func DeleteUserDesire(c *gin.Context) {
	DesireID := c.MustGet("wish_id").(int)
	res := model.DeleteWish(DesireID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 用户取消点亮他人愿望()
func CancelUserLight(c *gin.Context) {
	DesireID := c.MustGet("wish_id").(int)
	res := model.CancelLightDesire(DesireID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// 用户实现愿望
func AchieveUserDesire(c *gin.Context) {
	DesireID := c.MustGet("wish_id").(int)
	res := model.AchieveDesire(DesireID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}
