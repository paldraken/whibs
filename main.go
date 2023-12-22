package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/paldraken/whibs/configs"
	"github.com/paldraken/whibs/internal/parser"
	"github.com/paldraken/whibs/internal/server"
	"github.com/paldraken/whibs/internal/types"
	"github.com/paldraken/whibs/internal/watcher"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func main() {

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := "file://" + basePath + string(os.PathSeparator) + "client.html"
	fmt.Printf("You can start client in your browser by this link:\n\n%s\n\n", client)

	done := make(chan bool)
	lines := make(chan string, 10)
	rowCh := make(chan *types.SqlDebugRow)

	go watcher.WatchChanges(configs.SqlLogPath, done, lines)

	go parser.Lines(lines, rowCh)

	for row := range rowCh {
		server.NotifyWsUsers(row)
		row = nil
	}

	<-make(chan struct{})
}
