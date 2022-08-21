package model

import "test/common"

type User struct {
	ID           int           `json:"id" gorm:"id"`
	IdcardNumber string        `json:"idcard_number" gorm:"column:student_number,omitempty"`
	Password     string        `json:"password" gorm:"password,omitempty"`
	School       int           `json:"school" gorm:"school,omitempty"`
	Wechat       string        `json:"wechat" gorm:"wechat,omitempty"`
	Name         string        `json:"name" gorm:"name,omitempty"`
	Gender       common.Gender `json:"gender" gorm:"gender,omitempty"`
	Tel          string        `json:"tel" gorm:"tel,omitempty"`
	Email        string        `json:"email" gorm:"email,omitempty"`
	Major        string        `json:"major" gorm:"major,omitempty"`
	Desires      []Desire      `gorm:"foreignkey:UserID"`  //建立外键，为结构体Desire中的UserID
	Lights       []Desire      `gorm:"foreignkey:LightID"` //建立外键，为结构体Desire中的LightID
}

type ViewUser struct {
	Email     string `json:"email"`
	Wechat    string `json:"wechat"`
	Telephone string `json:"telephone"`
}

func CreateUser(data User) error {
	err := db.Create(&data).Error
	return err
}

func UpdateUser(user *User) error {
	err := db.Model(&User{}).Updates(user).Error
	return err
}

func GetUserIDByStudentNumber(studentNumber string) (int, error) {
	var user User
	err := db.Model(&User{}).Where("studentNumber = ?", studentNumber).First(&user).Error
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}

func GetUserInfo(UserID int) (*User, error) {
	var user *User
	err := db.Model(&User{}).Where("id = ?", UserID).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserCheck(email string) error {
	var user *User
	err := db.Model(&User{}).Where("email = ?", email).First(user).Error
	return err
}
