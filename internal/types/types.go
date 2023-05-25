package types

import (
	"strings"

	"github.com/paldraken/sqldebugwatch/internal/utils"
)

type SqlDebugRow struct {
	Time    float32  `json:"time"`
	Session string   `json:"session"`
	Sql     string   `json:"sql"`
	Trace   []string `json:"trace"`
	Modules []string `json:"modules"`
	Table   string   `json:"table"`
}

type Filter struct {
	Module    string `json:"module"`
	Table     string `json:"table"`
	WholeSql  string `json:"wholeSql"`
	Trace     string `json:"trace"`
	ShowTrace bool   `json:"showTrace"`
	Pause     bool   `json:"pause"`
}

func (f *Filter) IsPassed(row *SqlDebugRow) bool {

	if f.Pause {
		return false
	}

	if f.Module != "" && !utils.Contains(row.Modules, f.Module) {
		return false
	}

	if f.Table != "" && !strings.Contains(row.Table, f.Table) {
		return false
	}

	if f.WholeSql != "" && !strings.Contains(row.Sql, f.WholeSql) {
		return false
	}

	if f.Trace != "" {
		for _, t := range row.Trace {
			if strings.Contains(t, f.Trace) {
				return true
			}
		}
		return false
	}

	return true
}
