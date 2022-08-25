package model

type User struct {
	ID           int      `json:"id" gorm:"id"`
	IdcardNumber string   `json:"idcard_number" gorm:"column:student_number,omitempty"`
	Password     string   `json:"password" gorm:"password,omitempty"`
	School       int      `json:"school" gorm:"school,omitempty"`
	Wechat       string   `json:"wechat" gorm:"wechat,omitempty"`
	QQ           string   `json:"qq" gorm:"qq,omitempty"`
	Name         string   `json:"name" gorm:"name,omitempty"`
	Gender       int      `json:"gender" gorm:"gender,omitempty"`
	Tel          string   `json:"tel" gorm:"tel,omitempty"`
	Email        string   `json:"email" gorm:"email,omitempty"`
	Major        string   `json:"major" gorm:"major,omitempty"`
	Desires      []Desire `gorm:"foreignkey:UserID"`  //建立外键，为结构体Desire中的UserID
	Lights       []Desire `gorm:"foreignkey:LightID"` //建立外键，为结构体Desire中的LightID
}

type ViewUser struct {
	Name   string `json:"name"`
	QQ     string `json:"qq"`
	Email  string `json:"email"`
	Wechat string `json:"wechat"`
	Tel    string `json:"tel"`
}

type UserLogin struct {
	Email    string
	Password string
}

func CreateUser(user *User) error {
	err := db.Model(&User{}).Create(user).Error
	return err
}

func UpdateUser(user *User) error {
	err := db.Model(&User{}).Updates(user).Error
	return err
}

func GetUserIDByStudentNumber(studentNumber string) (int, error) {
	var user User
	err := db.Model(&User{}).Where("student_number = ?", studentNumber).First(&user).Error
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

func GetLightEmail(DesireID *int) (string, error) {
	var UserID int
	var email string
	err := db.Model(&Desire{}).Where("id = ?", DesireID).Select("light_id").Find(&UserID).Error
	if err != nil {
		return "", err
	}
	err = db.Model(&User{}).Where("id = ?", UserID).Select("email").Find(&email).Error
	return email, err
}

// // 通过UserID查询用户的邮箱是否存在
// func GetUserEmailByUserID(UserID int) string {
// 	var user User
// 	err := db.Model(&User{}).Where("id = ?", UserID).Find(&user).Error
// 	if err != nil {
// 		return ""
// 	}
// 	return user.Email
// }

// // 绑定邮箱
// func BindEmail(data User) helper.ReturnType {
// 	err := db.Model(&User{}).Where("student_number = ?", data.IdcardNumber).Updates(&data).Error
// 	if err != nil {
// 		return helper.ReturnType{Status: common.CodeError, Msg: "绑定邮箱失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: common.CodeSuccess, Msg: "绑定邮箱成功", Data: data.Email}
// 	}
// }
