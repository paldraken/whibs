package configs

type serverConfig struct {
	Enable bool
	Port   int
}

var ServerConfig serverConfig

func init() {
	initConfig()

	ServerConfig = serverConfig{
		Port:   k.Int("server.port"),
		Enable: k.Bool("server.enable"),
	}
}
