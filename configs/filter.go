package configs

import (
	"github.com/paldraken/sqldebugwatch/internal/types"
	"github.com/spf13/viper"
)

var ConsoleFilter types.Filter

func init() {
	initViper()
	ConsoleFilter = types.Filter{
		Module:    viper.GetString("console_filter.module"),
		Table:     viper.GetString("console_filter.table"),
		WholeSql:  "",
		Trace:     "",
		ShowTrace: viper.GetBool("console_filter.trace"),
		Pause:     viper.GetBool("console_filter.pause"),
	}
}
