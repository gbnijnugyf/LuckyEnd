package config

func GetSessionConfig() map[string]interface{} {
	sessionConfig := make(map[string]interface{})

	sessionConfig["key"] = "lucky"
	sessionConfig["name"] = "lucky_session"
	sessionConfig["age"] = 86400
	sessionConfig["path"] = "/"
	return sessionConfig
}
