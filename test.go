package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
var mysqlerr error
type Class struct {
	ID int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Students []Student `gorm:"foreignkey:ClassID"`
}

type Student struct {
	ID int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string
	ClassID int  `gorm:"class_id"`
}

func main(){
	dbDSN := "root:237156@(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	db, mysqlerr = gorm.Open("mysql", dbDSN)

	if mysqlerr != nil {
		panic("database open error "+ mysqlerr.Error())
	}

	var cla Class

	err := db.Model(&Class{}).Preload("Students").Where("ID = ?", 2).Find(&cla).Error
	if err != nil {
		panic("database find error "+ err.Error())
	}

	fmt.Println(cla)
}