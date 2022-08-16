package controller

import (
	"net/http"
	"strconv"
	"test/common"
	"test/helper"
	"test/model"
	"time"

	"github.com/gin-gonic/gin"
)

func UserAddDesire(c *gin.Context) {
	var desireJson model.Desire

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	UserID := c.MustGet("user_id").(int)
	desireJson.UserID = UserID
	// 检查用户当前许愿次数
	WishCount := model.GetUserDesireCount(&UserID)
	// 判断许愿总的次数是否超过上限
	if WishCount >= common.MaxWishCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "许愿次数已达上限", nil))
		return
	}

	desireJson.CreatAt = time.Now().In(common.ChinaTime)
	res := model.AddDesire(&desireJson)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "添加愿望失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "添加愿望成功", nil))
}

// UserLightDesire用户点亮他人愿望
func UserLightDesire(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	DesireId := c.PostForm("wish_id")
	DesireID, _ := strconv.Atoi(DesireId)
	LightCount := model.GetUserLightCount(&UserID)
	// 判断点亮次数是否达到上限
	if LightCount == common.GetCountError {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询错误", nil))
		return
	}
	if LightCount >= common.MaxLightCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "点亮愿望次数已达上限", nil))
		return
	}
	LightCount = model.GetUserLightMeantimeCount(&UserID)
	// 判断同时点亮次数是否达到上限
	if LightCount == common.GetCountError {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询错误", nil))
		return
	}
	if LightCount >= common.MaxLightSameCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "同时点亮愿望次数已达上限", nil))
		return
	}
	res := model.LightDesire(&DesireID, &UserID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

// GetUserAllDesires获取用户所有的愿望
func GetUserAllDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res, user := model.GetUserAllDesire(&UserID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", user))
}

// GetUserCreateDesires获取用户投递的愿望
func GetUserCreateDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res, desires := model.GetUserCreateDesire(&UserID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", desires))
}

// GetUserLightDesires获取用户点亮的愿望
func GetUserLightDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res, lights := model.GetUserLightDesire(&UserID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", lights))
}

// GetUserDesireByType通过点击分类查看愿望
func GetUserDesireByType(c *gin.Context) {
	Type := c.PostForm("type")
	desireType, _ := strconv.Atoi(Type)
	res, desires := model.GetDesireByCategories(&desireType)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", desires))
}

// DeleteUserDesire用户删除愿望
func DeleteUserDesire(c *gin.Context) {
	DesireID := c.MustGet("wish_id").(int)
	res := model.DeleteDesire(&DesireID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeSuccess, "删除失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))
}

// CancelUserLight用户取消点亮他人愿望
func CancelUserLight(c *gin.Context) {
	DesireID := c.MustGet("wish_id").(int)
	res := model.CancelLightDesire(&DesireID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "取消点亮失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "取消点亮成功", nil))
}

// AchieveUserDesire用户实现愿望
func AchieveUserDesire(c *gin.Context) {
	DesireID := c.MustGet("wish_id").(int)
	res := model.AchieveDesire(&DesireID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "实现失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "实现成功", nil))
}
