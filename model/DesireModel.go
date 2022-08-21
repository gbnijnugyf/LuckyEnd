package model

import (
	"fmt"
	"test/common"
	"test/helper"
	"time"
)

type ViewDesire struct {
	Desire   Desire   `json:"desire"`
	ViewUser ViewUser `json:"view_user"`
}

type Desire struct {
	ID       int       `json:"wish_id" gorm:"id" uri:"wish_id" form:"wish_id"`
	Desire   string    `json:"wish" gorm:"desire"`
	LightAt  time.Time `json:"light_at,omitempty" gorm:"light_at"`
	CreatAt  time.Time `json:"creat_at" gorm:"creat_at"`
	FinishAt time.Time `json:"finish_at" gorm:"finish_at"`
	State    int       `json:"state" gorm:"state"`
	Type     int       `json:"type" gorm:"type" form:"categories"`
	School   int       `json:"school" gorm:"school"`
	LightID  int       `json:"lightId" gorm:"light_id"` //点亮人外键
	UserID   int       `json:"userId" gorm:"user_id"`   //投递者外键
}

// AddDesire 用户添加愿望
func AddDesire(data *Desire) error {
	err := db.Model(&Desire{}).Omit("light_at").Create(data).Error
	return err
}

// LightDesire 用户点亮他人愿望
func LightDesire(DesireID *int, ID *int) helper.ReturnType {
	var desire Desire
	err := db.Model(&Desire{}).Where("id = ?", *DesireID).Find(&desire).Error
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
		desire.LightID = *ID
		err := db.Model(&Desire{}).Updates(&desire).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败，数据库发生错误", Data: err.Error()}
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "点亮愿望成功", Data: desire}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "点亮愿望失败", Data: desire}
}

// AchieveDesire 用户实现自己愿望
func AchieveDesire(DesireID *int) error {
	err := db.Model(&Desire{}).Where("id = ?", DesireID).Updates(Desire{State: common.WishHaveRealize, FinishAt: time.Now().In(common.ChinaTime)}).Error
	return err
}

// GetDesireByCategories 按分类查看愿望
func GetDesireByCategories(typ *int) (bool, []*Desire) {
	var desire []*Desire
	err := db.
		Select([]string{"id", "desire", "user_id", "creat_at", "light_at", "state"}).
		Where("type = ? AND	state = ?", *typ, common.WishNotLight).
		Find(&desire).
		Error
	if err != nil {
		return false, nil
	}
	return true, desire
}

// todo: 把所有的model加上error返回
func DeleteDesire(ID *int) bool {
	err := db.Model(&Desire{}).Where("user_id = ?", *ID).
		Update("state", common.WishHaveDelete).
		Error
	if err != nil {
		fmt.Println("delete desire error:" + err.Error())
		return false
	}
	return true
}

// 用户取消点亮愿望
func CancelLightDesire(DesireID *int) bool {
	err := db.
		Model(&Desire{}).Where("id = ?", *DesireID).
		Updates(Desire{State: common.WishNotLight, LightID: -1}).
		Error
	if err != nil {
		fmt.Println("cancel light error:" + err.Error())
		return false
	}
	return true
}

// 查询投递愿望数量
func GetUserDesireCount(ID *int) int64 {
	var count int64
	err := db.
		Model(&Desire{}).Where("user_id = ?", *ID).Count(&count).Error
	if err != nil {
		fmt.Println("get desire count error:" + err.Error())
		return common.GetCountError
	}
	return count
}

// 查询总共点亮愿望数量
func GetUserLightCount(ID *int) int64 {
	var count int64
	err := db.
		Model(&Desire{}).Where("light_id = ?", *ID).Count(&count).Error
	if err != nil {
		fmt.Println("get light count error:" + err.Error())
		return common.GetCountError
	}
	return count
}

// 查询同时点亮愿望数量
func GetUserLightMeantimeCount(ID *int) int64 {
	var count int64
	err := db.
		Model(&Desire{}).Where("light_id = ? AND state = ?", *ID, common.WishHaveLight).Count(&count).Error
	if err != nil {
		fmt.Println("get light meantime count error:" + err.Error())
		return common.GetCountError
	}
	return count
}
