package configs

import (
	"github.com/spf13/viper"
)

type serverConfig struct {
	Enable bool
	Port   int
}

var ServerConfig serverConfig

func init() {
	initViper()

	ServerConfig = serverConfig{
		Port:   serverPort(),
		Enable: isEnable(),
	}
}

func isEnable() bool {
	return viper.GetBool("server.enable")
}

func serverPort() int {
	return viper.GetInt("server.port")
}
