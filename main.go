package main

import (
	"log"

	"github.com/paldraken/whibs/configs"
	"github.com/paldraken/whibs/internal/parser"
	"github.com/paldraken/whibs/internal/server"
	"github.com/paldraken/whibs/internal/types"
	"github.com/paldraken/whibs/internal/watcher"
	"github.com/paldraken/whibs/internal/wconsole"
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
