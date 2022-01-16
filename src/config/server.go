package config

var serverConfig map[string]interface{}

func init() {
	serverConfig = make(map[string]interface{})

	serverConfig["host"] = ""
	serverConfig["port"] = "4001"

	serverConfig["mode"] = "debug"
}

func GetServerConfig() map[string]interface{} {
	return serverConfig
}
