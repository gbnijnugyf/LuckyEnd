package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	Username   string
	Password   string
	Hostname   string
	Port       string
	Dbname     string
	Charset    string
	PareseTime string
	Local      string
}

func GetDbConfig() DbConfig {
	var dbConfig DbConfig
	yamlFile, err := os.ReadFile("../conf/dbConfig.yaml")
	if err != nil {
		fmt.Println("read file error:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &dbConfig)
	if err != nil {
		fmt.Println("unmarshal error:" + err.Error())
	}
	return dbConfig
}
