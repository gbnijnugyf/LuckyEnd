package model

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
)

type ViewDesire struct {
	Desire   Desire   `json:"desire"`
	ViewUser ViewUser `json:"view_user"`
}

type Desire struct {
	ID        int       `json:"desire_id" gorm:"id" uri:"desire_id" form:"desire_id"`
	Desire    string    `json:"desire" gorm:"desire"`
	LightAt   time.Time `json:"light_at,omitempty" gorm:"light_at"`
	CreatAt   time.Time `json:"creat_at" gorm:"creat_at"`
	FinishAt  time.Time `json:"finish_at" gorm:"finish_at"`
	State     int       `json:"state" gorm:"state"`
	Type      int       `json:"type" gorm:"type" form:"categories"`
	School    int       `json:"school" gorm:"school"`
	LightID   int       `gorm:"light_id"` //点亮人外键
	UserID    int       `gorm:"user_id"`  //投递者外键
	LightUser ViewUser  `gorm:"ForeginKey:LightID;AssociationForeignKey:ID"`
}

func AddDesire(data *Desire) error {
	err := db.Model(&Desire{}).Omit("light_at").Create(data).Error
	return err
}

func LightDesire(DesireID *int, ID *int) helper.ReturnType {
	var desire Desire
	err := db.Model(&Desire{}).Where("id = ?", *DesireID).Find(&desire).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败,未能查询到该愿望", Data: err.Error()}
	}
	if desire.State == common.DesireHaveLight {
		return helper.ReturnType{Status: common.CodeError, Msg: "点亮愿望失败，该愿望已经被人抢先点亮了", Data: nil}
	}
	// 愿望处于未点亮状态则点亮
	if desire.State == common.DesireNotLight {
		desire.State = common.DesireHaveLight
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

func AchieveDesire(DesireID *int) error {
	err := db.Model(&Desire{}).Where("id = ?", DesireID).Updates(Desire{State: common.DesireHaveRealize, FinishAt: time.Now().In(common.ChinaTime)}).Error
	return err
}

func GetDesireByCategories(typ *int) (bool, []*Desire) {
	var desire []*Desire
	err := db.
		Select([]string{"id", "desire", "user_id", "creat_at", "light_at", "state"}).
		Where("type = ? AND	state = ?", *typ, common.DesireNotLight).
		Find(&desire).
		Error
	if err != nil {
		log.Errorf("Error in get desireBycatgories: %+v", errors.WithStack(err))
		return false, nil
	}
	return true, desire
}

func DeleteDesire(ID *int) error {
	err := db.Model(&Desire{}).Where("user_id = ?", *ID).Update("state", common.DesireHaveDelete).Error
	return err
}

func CancelLightDesire(DesireID *int) error {
	err := db.Model(&Desire{}).Where("id = ?", *DesireID).Updates(Desire{State: common.DesireNotLight, LightID: -1}).Error
	return err
}

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

func GetUserLightMeantimeCount(ID int) int64 {
	var count int64
	err := db.
		Model(&Desire{}).Where("light_id = ? AND state = ?", ID, common.DesireHaveLight).Count(&count).Error
	if err != nil {
		log.Errorf("get light meantime count error:%+v", errors.WithStack(err))
		return common.GetCountError
	}
	return count
}

func DesireContentByID(DesireID *int) (string, error) {
	var desire Desire
	err := db.Model(&Desire{}).Where("id = ?", *DesireID).Find(&desire).Error
	return desire.Desire, err
}

func GetDesire(DesireID *int) (Desire, error) {
	var desire Desire
	err := db.Model(&Desire{}).Where("id = ?", *DesireID).Find(&desire).Error
	return desire, err
}
