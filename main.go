package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("failed to create the new watcher: %v", err)
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {

			select {

			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op == fsnotify.Write { //== fsnotify.Write
					fmt.Println(time.Now().Format("01-02-2006 15:04:05"))
					log.Println("modified file:", event.Name)
					printContent("config_dir")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)

			}
		}
	}()

	err = watcher.Add("config_dir")
	if err != nil {
		log.Printf("failed to add directory for watch: %v", err)
		os.Exit(1)
	}
	log.Println(watcher.WatchList())
	<-done
}

func printContent(name string) {
	fileList, err := ioutil.ReadDir(name)
	if err != nil {
		log.Printf("failed to read the directory: %v", err)
	}

	for _, fileName := range fileList {
		if !fileName.IsDir() {
			byteData, err := ioutil.ReadFile(name + "/" + fileName.Name())
			if err != nil {
				log.Printf("failed to read the  data from specified file: %v", err)
			}

			log.Println(string(byteData))
		}

	}
}
