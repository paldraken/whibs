package configs

import (
	"fmt"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var once sync.Once

func initViper() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		// pflag.Bool("se", false, "Enable ws server")
		pflag.String("m", "", "module name")
		pflag.String("p", "", "path to sql log")
		pflag.Parse()

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		// command line
		viper.BindPFlags(pflag.CommandLine)
		viper.RegisterAlias("path", "p")
		viper.RegisterAlias("console_filter.module", "m")
		// viper.RegisterAlias("server.enable", "se")

		fmt.Println("enable", viper.GetString("path"))
	})
}
