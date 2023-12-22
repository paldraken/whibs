package configs

import (
	"github.com/paldraken/whibs/internal/types"
)

var ConsoleFilter types.Filter

func init() {
	initConfig()

	ConsoleFilter = types.Filter{
		Module:    k.String("filter.module"),
		Table:     k.String("filter.table"),
		WholeSql:  k.String("filter.sql"),
		Trace:     k.String("filter.trace"),
		ShowTrace: k.Bool("filter.show_trace"),
		Pause:     k.Bool("filter.pause"),
	}
}
