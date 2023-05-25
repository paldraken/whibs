package wconsole

import (
	"fmt"

	"github.com/paldraken/sqldebugwatch/configs"
	"github.com/paldraken/sqldebugwatch/internal/types"
)

func Print(row *types.SqlDebugRow) {

	if !configs.ConsoleFilter.IsPassed(row) {
		return
	}

	printDebugRow(row)

}

func printDebugRow(row *types.SqlDebugRow) {
	s := `
Time: %f, Session: %s, 	

%s 
	
%v 
Modules: %v
Table: %s
Modules: %v

*************************`

	var trace []string
	if configs.ConsoleFilter.ShowTrace {
		trace = row.Trace
	}

	fmt.Printf(s, row.Time, row.Session, row.Sql, trace, row.Modules, row.Table, row.Modules)
}
