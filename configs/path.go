package configs

import (
	"log"
)

var SqlLogPath string

func init() {
	initConfig()

	SqlLogPath = k.String("path")

	if SqlLogPath == "" {
		log.Fatal("path to sql log is required")
	}
}
