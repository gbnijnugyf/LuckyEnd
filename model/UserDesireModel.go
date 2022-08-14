package model

import (
	"fmt"
)

// 查询用户的所有愿望，包括投递和点亮的愿望
func GetUserAllDesire(UserID *int) (bool, User) {
	var user User
	err := db.Model(&User{}).Preload("desires", "lights").Where("id = ?", *UserID).Find(&user).Error
	if err != nil {
		fmt.Println("get desier error:" + err.Error())
		return false, user
	}
	return true, user
}

// 查询用户的投递愿望
func GetUserCreateDesire(UserID *int) (bool, []*Desire) {
	var desires []*Desire
	err := db.Model(&Desire{}).Where("user_id = ?", UserID).Find(&desires).Error
	if err != nil {
		fmt.Println("get desire error:" + err.Error())
		return false, nil
	}
	return true, desires
}

// 查询用户的点亮愿望
func GetUserLightDesire(UserID *int) (bool, []*Desire) {
	var lights []*Desire
	err := db.Model(&Desire{}).Where("user_id = ?", *UserID).Find(&lights).Error
	if err != nil {
		return false, nil
	}
	return true, lights
}
