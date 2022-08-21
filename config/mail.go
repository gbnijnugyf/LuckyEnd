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
	Username string
	Password string
}

func GetMailConfig() MailConfig {
	var mailconfig MailConfig
	yamlFile, err := os.ReadFile("./mailconfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &mailconfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return mailconfig
}
