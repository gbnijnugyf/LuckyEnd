package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type mailConfig struct {
	Host     string
	Port     int
	From     string
	Username string
	Password string
}

func GetMailConfig() mailConfig {
	var mailconfig mailConfig
	yamlFile, err := ioutil.ReadFile("D:/yaml/mailconfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &mailconfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return mailconfig
}
