package controller

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UserAddDesire(c *gin.Context) {
	desireFormView := &model.ViewDesire{}

	if err := c.ShouldBindJSON(desireFormView); err != nil {
		log.Errorf("request param error %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err))
		return
	}
	desire := desireFormView.Desire
	UserID := c.MustGet("user_id").(int)
	desire.UserID = UserID
	desire.State = common.DesireNotLight
	// 检查用户当前许愿次数
	WishCount := model.GetUserDesireCount(&UserID)
	// 判断许愿总的次数是否超过上限
	if WishCount >= common.MaxDesireCount {
		log.Errorf("you have to many desires %+v", errors.New(helper.ErrTooManyDesires))
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "许愿次数已达上限", nil))
		return
	}
	err := model.AddDesire(&desire)
	if err != nil {
		log.Errorf("Failed to add desire error is :%+v", err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "添加愿望失败", err))
		return
	}
	err = model.UpdateUser(&model.User{
		ID:     UserID,
		Name:   desireFormView.ViewUser.Name,
		QQ:     desireFormView.ViewUser.QQ,
		Email:  desireFormView.ViewUser.Email,
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
	lightFromView := &model.ViewLight{}
	if err := c.ShouldBindJSON(lightFromView); err != nil {
		log.Errorf("request param error %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err))
		return
	}
	UserID := c.MustGet("user_id").(int)
	DesireID := lightFromView.DesireID
	// if err != nil {
	// 	log.Errorf("Wrong desireId err:%+v", errors.WithStack(err))
	// 	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "bad params", err))
	// 	return
	// }
	checkID, err := model.GetUserID(&DesireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询错误", err))
		return
	}
	if checkID == UserID {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "无法点亮自己的愿望哦", nil))
		return
	}
	LightCount := model.GetUserLightCount(&UserID)
	// 判断点亮次数是否达到上限
	if LightCount == common.GetCountError {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "查询错误", nil))
		return
	}
	if LightCount >= common.MaxLightCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "点亮愿望次数已达上限", nil))
		return
	}
	LightCount = model.GetUserLightMeantimeCount(UserID)
	// 判断同时点亮次数是否达到上限
	if LightCount == common.GetCountError {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询错误", nil))
		return
	}
	if LightCount >= common.MaxLightSameCount {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "同时点亮愿望次数已达上限", nil))
		return
	}
	res := model.LightDesire(&DesireID, &UserID)
	if res.Status == common.CodeError {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	email, err := model.GetEmail(&DesireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "发送邮件失败", err))
		return
	}
	go func() {
		err := helper.SendMail(email, common.LightDesire, "", "")
		if err != nil {
			log.Print(err)
		}
	}()
	err = model.UpdateUser(&model.User{
		ID:     UserID,
		Name:   lightFromView.Name,
		QQ:     lightFromView.QQ,
		Wechat: lightFromView.Wechat,
		Tel:    lightFromView.Tel,
	})
	if err != nil {
		log.Errorf("Failed to update user err: %+v", err)
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserCreateDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	desires, err := model.GetUserCreateDesire(&UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", desires))
}

func GetUserLightDesires(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	lights, err := model.GetUserLightDesire(&UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询失败", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", lights))
}

func GetUserDesireByType(c *gin.Context) {
	Type, ok := c.GetQuery("type")
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
	desires, err := model.GetDesireByCategories(&desireType)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "查询成功", desires))
}

func DeleteUserDesire(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	DesireID := c.Query("desire_id")
	desireID, err := strconv.Atoi(DesireID)
	if err != nil {
		log.Errorf("cannot convert desireid into int err :%+v", err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "desire_id is not right", err))
		return
	}
	checkID, err := model.GetUserID(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除失败", err))
		return
	}
	if UserID != checkID {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "不能删除不是你的愿望哦", nil))
		return
	}
	email, err := model.GetLightEmail(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "邮件发送失败", err))
		return
	}
	desireContent, err := model.DesireContentByID(&desireID)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "邮件发送失败", err))
		return
	}
	res := model.DeleteDesire(&desireID)
	if res != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "删除失败", res))
		return
	}
	go func() {
		err := helper.SendMail(email, common.DeleteDesire, desireContent, "")
		if err != nil {
			log.Print(err)
		}
	}()
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))
}

func CancelUserLight(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	json := make(map[string]interface{})
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Errorf("request param error %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "绑定数据失败", err))
		return
	}
	desireID := int(json["desire_id"].(float64))
	message := json["message"].(string)
	checkID, err := model.GetLightID(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "取消点亮失败,无法验证身份", nil))
		return
	}
	if UserID != checkID {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "无法取消不是你点亮的愿望哦", nil))
		return
	}
	desireContent, err := model.DesireContentByID(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "取消点亮失败", err))
		return
	}
	email, err := model.GetEmail(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "取消点亮失败", err))
		return
	}
	go func() {
		err := helper.SendMail(email, common.CancelLight, desireContent, message)
		if err != nil {
			log.Print(err)
		}
	}()
	res := model.CancelLightDesire(&desireID)
	if res != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "取消点亮失败", res))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "取消点亮成功", nil))
}

func AchieveUserDesire(c *gin.Context) {
	UserID := c.MustGet("user_id").(int)
	json := make(map[string]interface{})
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Errorf("request param error %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "绑定数据失败", err))
		return
	}
	desireID := int(json["desire_id"].(float64))
	checkID, err := model.GetUserID(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "实现愿望失败", err))
		return
	}
	if UserID != checkID {
		email, err := model.GetEmail(&desireID)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "邮件发送失败", err))
			return
		}
		name, err := model.GetName(UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "邮件发送失败", err))
			return
		}
		go func() {
			err := helper.SendMail(email, common.HaveAchieve, name, "")
			if err != nil {
				log.Print(err)
			}
		}()
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "邮件发送成功,我们已经通知该同学啦", nil))
		return
	}
	res := model.AchieveDesire(&desireID)
	if res != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "实现失败", res))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "实现成功", nil))
}

func DesireDetail(c *gin.Context) {
	DesireID := c.Query("desire_id")
	desireID, _ := strconv.Atoi(DesireID)
	desire, err := model.GetInfo(&desireID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取信息失败", err))
		return
	}
	user, err := model.GetUserInfo(desire.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取信息失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取信息成功", &model.ViewDesire{
		Desire: *desire,
		ViewUser: model.ViewUser{
			Name:   user.Name,
			QQ:     user.QQ,
			Email:  user.Email,
			Wechat: user.Wechat,
			Tel:    user.Tel,
			School: user.School,
		},
	}))
}
