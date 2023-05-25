package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var SqlLogPath string

func init() {
	initViper()

	SqlLogPath = viper.GetString("path")

	fmt.Println("path", SqlLogPath)

	if SqlLogPath == "" {
		log.Fatal("path to sql log is empty")
	}
}
