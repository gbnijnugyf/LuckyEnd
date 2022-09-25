package config

import (
	"os"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	FromName string `yaml:"fromName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetMailConfig() MailConfig {
	var mailConfig MailConfig
	yamlFile, err := os.ReadFile("./conf/mailConfig.yaml")
	if err != nil {
		log.Errorf("read file error + %+v", errors.WithStack(err))
	}
	err = yaml.Unmarshal(yamlFile, &mailConfig)
	if err != nil {
		log.Errorf("unmarshal error + %+v", errors.WithStack(err))
	}
	return mailConfig
}
