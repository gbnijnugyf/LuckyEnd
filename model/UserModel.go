package model

import (
	"test/helper"
	"test/common"
)

type User struct {
	ID           int    `json:"id" gorm:"id"`
	IdcardNumber string `json:"idcard_number" gorm:"column:student_number"`
	Password     string `json:"password" gorm:"password"`
	School       int    `json:"school" gorm:"school"`
	Wechat       string `json:"wechat" gorm:"wechat"`
	Name         string `json:"name" gorm:"name"`
	Gender       int    `json:"gender" gorm:"gender"`
	Tel          string `json:"tel" gorm:"tel"`
	Email        string `json:"email" gorm:"email"`
	Major        string `json:"major" gorm:"major"`
	Desires	     []Desire `gorm:"foreignkey:UserID"`	//建立外键，为结构体Desire中的UserID
	Lights       []Desire `gorm:"foreignkey:LightID"`	//建立外键，为结构体Desire中的LightID
}

// 登录检查
func (model *User) LoginCheck(data User) helper.ReturnType {
	user := User{}
	err := db.Where("student_number = ? AND password = ?", data.IdcardNumber, data.Password).First(&user).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "登录验证成功", Data: user}
	}
}

// 绑定邮箱
func (model *User) BindEmail(data User) helper.ReturnType {
	err := db.Model(&User{}).Where("student_number = ?", data.IdcardNumber).Update(&data).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "绑定邮箱失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "绑定邮箱成功", Data: data.Email}
	}
}

// 创建新用户
func (model *User) CreateUser(data User) helper.ReturnType {
	err := db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "创建新用户失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "创建新用户成功", Data: data}
}
