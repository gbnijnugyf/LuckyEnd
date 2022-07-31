package model

import (
	"log"
	"test/common"
	"test/helper"
	"time"
)

type Desire struct {
	ID       int       `json:"wish_id" gorm:"id" uri:"wish_id" form:"wish_id"`
	Desire   string    `json:"wish" gorm:"desire"`
	LightAt  time.Time `json:"light_at,omitempty" gorm:"light_at"`
	CreatAt  time.Time `json:"creat_at" gorm:"creat_at"`
	FinishAt time.Time `json:"finish_at" gorm:"finsh_at"`
	State    int       `json:"state" gorm:"state"`
	Type     int       `json:"type" gorm:"type" form:"categories"`
	School   int       `json:"school" gorm:"school"`
	LightID  int       `gorm:"light_id"` //点亮人外键
	UserID   int       `gorm:"user_id"`  //投递者外键
}

// 用户添加愿望
func AddDesire(data Desire) helper.ReturnType {
	log.Println(data.CreatAt)
	err := db.Model(&Desire{}).Omit("light_at").Create(&data).Error // 没有.Error会报错

	if err != nil {
		log.Print("43", err)
		return helper.ReturnType{Status: common.CodeError, Msg: "添加愿望失败", Data: err}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加愿望成功", Data: data}
}

// 用户点亮他人愿望
func LightDesire(DesireID int, ID int) helper.ReturnType {
	var desire Desire
	err := db.Model(&Desire{}).Where("id = ?", DesireID).Find(&desire).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败,未能查询到该愿望", Data: err.Error()}
	}
	if desire.State == common.WishHaveLight {
		return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败，该愿望已经被人抢先点亮了", Data: nil}
	}
	// 愿望处于未点亮状态则点亮
	if desire.State == common.WishNotLight {
		desire.State = common.WishHaveLight
		desire.LightAt = time.Now().In(common.ChinaTime)
		desire.LightID = ID
		err := db.Model(&Desire{}).Update(&desire).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败，数据库发生错误", Data: err.Error()}
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "点亮愿望成功", Data: desire}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "点亮愿望失败", Data: desire}
}

// 用户实现自己愿望
func AchieveDesire(DesireID int) helper.ReturnType {
	var desire Desire
	err := db.Model(&Desire{}).Where("id = ?", DesireID).Update("state", common.WishHaveRealize, "finish_at", time.Now().In(common.ChinaTime)).Find(&desire).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望成功", Data: desire}
}

// 按分类查看愿望
func GetWishByCategories(typ int) helper.ReturnType {
	var desire []*Desire
	err := db.
		Select([]string{"id", "desire", "user_id", "creat_at", "light_at", "state"}).
		Where("type = ? AND	state = ?", typ, common.WishNotLight).
		Find(&desire).
		Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查看愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查看愿望成功", Data: desire}
}

// 用户删除愿望
func DeleteWish(ID int) helper.ReturnType {
	err := db.Model(&Desire{}).Where("user_id = ?", ID).
		Update("state", common.WishHaveDelete).
		Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除愿望成功", Data: ""}
}

// 用户取消点亮愿望
func CancelLightDesire(DesireID int) helper.ReturnType {
	err := db.
		Model(&Desire{}).Where("id = ?", DesireID).
		Update(map[string]interface{}{"state": common.WishNotLight, "light_id": -1}).
		Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "取消点亮愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "取消点亮愿望成功", Data: nil}
}

// 查询投递愿望数量
func GetUserWishCount(ID int) int {
	var count int
	err := db.
		Model(&Desire{}).Where("user_id = ?", ID).Count(&count).Error
	if err != nil {
		return -1
	}
	return count
}

// 查询总共点亮愿望数量
func GetUserLightCount(ID int) int {
	var count int
	err := db.
		Model(&Desire{}).Where("light_id = ?", ID).Count(&count).Error
	if err != nil {
		return -1
	}
	return count
}

// 查询同时点亮愿望数量
func GetUserLightAtSameTimeCount(ID int) int {
	var count int
	err := db.
		Model(&Desire{}).Where("light_id = ? AND state = ?", ID, common.WishHaveLight).Count(&count).Error
	if err != nil {
		return -1
	}
	return count
}
