package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Hostname   string `yaml:"hostname"`
	Port       string `yaml:"port"`
	Dbname     string `yaml:"dbname"`
	Charset    string `yaml:"charset"`
	PareseTime string `yaml:"paresetime"`
	Local      string `yaml:"local"`
}

func GetDbConfig() DbConfig {
	var dbConfig DbConfig
	yamlFile, err := os.ReadFile("D:/golang/conf/dbConfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &dbConfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return dbConfig
}
