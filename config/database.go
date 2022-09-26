package config

import (
	"os"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type dbConfig struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Hostname  string `yaml:"hostname"`
	Port      string `yaml:"port"`
	Dbname    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Local     string `yaml:"local"`
}

type DbConfig struct {
	DbConfig dbConfig `yaml:"dbConfig"`
}

func GetDbConfig() dbConfig {
	var dbConfig DbConfig
	yamlFile, err := os.ReadFile("./conf/dbConfig.yaml")
	if err != nil {
		log.Errorf("read file error: %+v", errors.WithStack(err))
	}
	err = yaml.Unmarshal(yamlFile, &dbConfig)
	if err != nil {
		log.Errorf("unmarshal error: %+v", errors.WithStack(err))
	}
	return dbConfig.DbConfig
}
