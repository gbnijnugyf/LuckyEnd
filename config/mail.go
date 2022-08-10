package config

var mailConfig map[string]interface{}

func init() {
	mailConfig = make(map[string]interface{})
	mailConfig["host"] = "smtp.163.com"
	mailConfig["port"] = 465
	mailConfig["from"] = "goproject"
	mailConfig["username"] = "luxuetao0518@163.com"
	mailConfig["password"] = "CLSVNTLRRSTGAZTY"
}

func GetMailConfig() map[string]interface{} {
	return mailConfig
}