package watcher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/paldraken/sqldebugwatch/configs"
)

func WatchChanges(path string, done chan bool, lines chan string) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {

					scanner := bufio.NewScanner(file)
					scanner.Split(bufio.ScanLines)
					for scanner.Scan() {
						lines <- scanner.Text()
					}

					currentSize, err := file.Seek(0, io.SeekEnd)
					if err != nil {
						currentSize = 0
					}

					if isShrinkedOutsize(file, currentSize) {
						_, err = file.Seek(0, io.SeekEnd)
						if err != nil {
							fmt.Println(err)
							return
						}
					}

					// trunkate file
					if configs.ShrinkLog != nil && configs.ShrinkLog.Cmp(big.NewInt(currentSize)) < 0 {
						err := os.Truncate(configs.SqlLogPath, 0)
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func isShrinkedOutsize(file *os.File, current int64) bool {
	stat, err := file.Stat()
	if err != nil {
		return false
	}
	return stat.Size() < current
}
