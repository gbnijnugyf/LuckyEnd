package model

import (
	"test/common"
	"test/helper"
)

// 查询用户的所有愿望，包括投递和点亮的愿望
func GetUserAllDesire(UserID int) helper.ReturnType {
	var user User
	err := db.Model(&User{}).Preload("desires", "lights").Where("id = ?", UserID).Find(&user).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询所有愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询所有愿望成功", Data: user}
}

// 查询用户的投递愿望
func GetUserCreateDesire(UserID int) helper.ReturnType {
	var desires []Desire
	err := db.Model(&Desire{}).Where("user_id = ?", UserID).Find(&desires).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询投递愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询投递愿望成功", Data: err.Error()}
}

// 查询用户的点亮愿望
func GetUserLightDesire(UserID int) helper.ReturnType {
	var lights []Desire
	err := db.Model(&Desire{}).Where("user_id = ?", UserID).Find(&lights).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询点亮愿望失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询点亮愿望成功", Data: err.Error()}	
}