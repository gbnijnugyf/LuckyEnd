package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MailConfig struct {
	Host     string
	Port     int
	From     string
	Fromname string
	Username string
	Password string
}

func GetMailConfig() MailConfig {
	var mailConfig MailConfig
	yamlFile, err := ioutil.ReadFile("D:/yaml/mailconfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &mailConfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return mailConfig
}
