package model

import (
	"test/common"
	"test/helper"
)

type User struct {
	ID           int      `json:"id" gorm:"id"`
	IdcardNumber string   `json:"idcard_number" gorm:"column:student_number"`
	Password     string   `json:"password" gorm:"password"`
	School       int      `json:"school" gorm:"school"`
	Wechat       string   `json:"wechat" gorm:"wechat"`
	Name         string   `json:"name" gorm:"name"`
	Gender       int      `json:"gender" gorm:"gender"`
	Tel          string   `json:"tel" gorm:"tel"`
	Email        string   `json:"email" gorm:"email"`
	Major        string   `json:"major" gorm:"major"`
	Desires      []Desire `gorm:"foreignkey:UserID"`  //建立外键，为结构体Desire中的UserID
	Lights       []Desire `gorm:"foreignkey:LightID"` //建立外键，为结构体Desire中的LightID
}

// 登录检查
func LoginCheck(data User) helper.ReturnType {
	user := User{}
	err := db.Where("student_number = ? AND password = ?", data.IdcardNumber, data.Password).First(&user).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "登录验证成功", Data: user}
	}
}

// 绑定邮箱
func BindEmail(data User) helper.ReturnType {
	err := db.Model(&User{}).Where("student_number = ?", data.IdcardNumber).Updates(&data).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "绑定邮箱失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "绑定邮箱成功", Data: data.Email}
	}
}

// 创建新用户
func CreateUser(data User) helper.ReturnType {
	err := db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "创建新用户失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "创建新用户成功", Data: data}
}

// 通过student_number查询用户的ID
func GetUserIDByStudentNumber(student_number string) helper.ReturnType {
	var user User
	err := db.Model(&User{}).Where("student_number = ?", student_number).First(&user).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败，数据库错误", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: user}
}

// 通过UserID查询用户的邮箱是否存在
func GetUserEmailByUserID(UserID int) string {
	var user User
	err := db.Model(&User{}).Where("id = ?", UserID).Find(&user).Error
	if err != nil {
		return ""
	}
	return user.Email
}
