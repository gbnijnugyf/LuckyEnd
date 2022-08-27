package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type MailConfig struct {
	Host     string
	Port     int
	From     string
	FromName string
	Username string
	Password string
}

func GetMailConfig() MailConfig {
	var mailConfig MailConfig
	yamlFile, err := os.ReadFile("../conf/mailConfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &mailConfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return mailConfig
}
