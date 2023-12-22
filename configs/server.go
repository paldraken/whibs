package configs

type serverConfig struct {
	Port int
}

var ServerConfig serverConfig

func init() {
	initConfig()

	ServerConfig = serverConfig{
		Port: k.Int("server.port"),
	}
}
