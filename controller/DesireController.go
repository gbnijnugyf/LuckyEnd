package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UserAddDesire(c *gin.Context) {
	var desireFormView model.ViewDesire

	if err := c.ShouldBindJSON(&desireFormView); err != nil {
		log.Errorf("request param error %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err))
		return
	}
	desire := desireFormView.Desire
	UserID := c.MustGet("user_id").(int)
	desireFormView.Desire.UserID = UserID
	// 检查用户当前许愿次数
	WishCount := model.GetUserDesireCount(&UserID)
	// 判断许愿总的次数是否超过上限
	if WishCount >= common.MaxWishCount {
		log.Errorf("you have to many desires %+v", errors.New(helper.ErrTooManyDesires))
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "许愿次数已达上限", nil))
		return
	}
	// todo: github.com/shawu21/test if it will give auto value
	desireFormView.Desire.CreatAt = time.Now().In(common.ChinaTime)
	err := model.AddDesire(&desire)
	if err != nil {
		log.Errorf("Failed to add desire error is :%+v", err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "添加愿望失败", err))
		return
	}
	err = model.UpdateUser(&model.User{
		ID:     UserID,
		Wechat: desireFormView.ViewUser.Wechat,
		Tel:    desireFormView.ViewUser.Tel,
	})
	// Here is no need to return bad request, because is an additional function
	if err != nil {
		log.Errorf("Failed to update user err: %+v", err)
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "添加愿望成功", nil))
}

func UserLightDesire(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	DesireId := c.PostForm("desire_id")
	DesireID, err := strconv.Atoi(DesireId)
	if err != nil {
		log.Errorf("Wrong desireId err:%+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "bad params", nil))
		return
	}
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
	// TODO: send email
	if res.Status == common.CodeError {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserAllDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res, user := model.GetUserAllDesire(&UserID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", user))
}

func GetUserCreateDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res, desires := model.GetUserCreateDesire(&UserID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", desires))
}

func GetUserLightDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	res, lights := model.GetUserLightDesire(&UserID)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", lights))
}

func GetUserDesireByType(c *gin.Context) {
	Type, ok := c.GetQuery("categoires")
	if !ok {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "type is empty", nil))
		return
	}
	desireType, err := strconv.Atoi(Type)
	if err != nil {
		log.Errorf("cannot convert type into int err :%+v", err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "type is not right", err))
		return
	}
	res, desires := model.GetDesireByCategories(&desireType)
	if !res {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", desires))
}

func DeleteUserDesire(c *gin.Context) {
	DesireID := c.PostForm("desire_id")
	desireID, _ := strconv.Atoi(DesireID)
	res := model.DeleteDesire(&desireID)
	if res != nil {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeSuccess, "删除失败", res))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))
}

func CancelUserLight(c *gin.Context) {
	DesireID := c.PostForm("desire_id")
	desireID, _ := strconv.Atoi(DesireID)
	message := c.PostForm("message")
	desireContent, err := model.DesireContentByID(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "取消点亮失败", err))
	}
	email, err := model.GetLightEmail(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "取消点亮失败", err))
	}
	go func() {
		err := helper.SendMail(email, common.CancelLight, desireContent, message)
		if err != nil {
			log.Print(err)
		}
	}()
	res := model.CancelLightDesire(&desireID)
	if res != nil {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "取消点亮失败", res))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "取消点亮成功", nil))
}

func AchieveUserDesire(c *gin.Context) {
	DesireID := c.PostForm("desire_id")
	desireID, _ := strconv.Atoi(DesireID)
	res := model.AchieveDesire(&desireID)
	if res != nil {
		c.JSON(http.StatusInternalServerError, helper.ApiReturn(common.CodeError, "实现失败", res))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "实现成功", nil))
}
