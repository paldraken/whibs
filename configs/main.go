package configs

import (
	"flag"
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

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		// command line
		viper.RegisterAlias("path", "p")
		flag.String("p", "", "path to sql log")

		viper.RegisterAlias("console_filter.module", "m")
		flag.String("m", "", "module name")

		viper.RegisterAlias("server.enable", "s")
		flag.String("s", "", "Enable ws server")

		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
		pflag.Parse()
		viper.BindPFlags(pflag.CommandLine)
	})
}
