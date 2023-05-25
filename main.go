package main

import (
	"log"

	"github.com/paldraken/sqldebugwatch/configs"
	"github.com/paldraken/sqldebugwatch/internal/parser"
	"github.com/paldraken/sqldebugwatch/internal/server"
	"github.com/paldraken/sqldebugwatch/internal/types"
	"github.com/paldraken/sqldebugwatch/internal/watcher"
	"github.com/paldraken/sqldebugwatch/internal/wconsole"
)

func main() {

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan bool)
	lines := make(chan string, 10)
	rowCh := make(chan *types.SqlDebugRow)

	go watcher.WatchChanges(configs.SqlLogPath, done, lines)

	go parser.Lines(lines, rowCh)

	for row := range rowCh {
		if !configs.ConsoleFilter.Pause {
			wconsole.Print(row)
		}
		server.NotifyWsUsers(row)
		row = nil
	}

	<-make(chan struct{})
}
