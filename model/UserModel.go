package model

// todo: 更改外键
type User struct {
	ID           int      `json:"id" gorm:"id;primary_key;auto_increment"`
	IdcardNumber string   `json:"idcard_number" gorm:"student_number,omitempty"`
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
	Lights       []Desire `gorm:"foreignkey:LightID"` //建立外键，为结构体Desire中的UserLight
}

type ViewUser struct {
	Name   string `json:"name"`
	QQ     string `json:"qq"`
	Email  string `json:"email"`
	Wechat string `json:"wechat"`
	Tel    string `json:"tel"`
}

type ViewLight struct {
	DesireID int    `json:"desire_id"`
	Name     string `json:"name"`
	QQ       string `json:"qq"`
	Wechat   string `json:"wechat"`
	Tel      string `json:"tel"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	err := db.Model(&User{}).Create(user).Error
	return err
}

func UpdateUser(user *User) error {
	err := db.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error
	return err
}

func GetUserIDByStudentNumber(studentNumber string) (int, error) {
	user := &User{}
	err := db.Model(&User{}).Where("idcard_number = ?", studentNumber).First(user).Error
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}

func GetUserInfo(UserID int) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("id = ?", UserID).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserCheck check user if exists
func UserCheck(email string) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("email = ?", email).First(user).Error
	return user, err
}

func GetLightEmail(DesireID *int) (string, error) {
	desire := &Desire{}
	user := &User{}
	err := db.Model(&Desire{}).Where("id = ?", DesireID).Select("light_id").Find(desire).Error
	if err != nil {
		return "", err
	}
	err = db.Model(&User{}).Where("id = ?", desire.LightID).Select("email").Find(user).Error
	return user.Email, err
}

func GetEmail(DesireID *int) (string, error) {
	desire := &Desire{}
	user := &User{}
	err := db.Model(&Desire{}).Where("id = ?", *DesireID).Select("user_id").Find(desire).Error
	if err != nil {
		return "", err
	}
	err = db.Model(&User{}).Where("id = ?", desire.UserID).Select("email").Find(user).Error
	return user.Email, err
}

func GetName(UserID *int) (string, error) {
	user := &User{}
	err := db.Model(&User{}).Where("id = ?", *UserID).Select("name").Find(user).Error
	return user.Name, err
}