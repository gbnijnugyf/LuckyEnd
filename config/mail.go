package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	FromName string `yaml:"fromname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetMailConfig() MailConfig {
	var mailConfig MailConfig
	yamlFile, err := os.ReadFile("D:/golang/conf/mailConfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &mailConfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return mailConfig
}
