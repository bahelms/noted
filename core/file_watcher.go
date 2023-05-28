package core

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/bahelms/noted/config"
)

type fileWatcher struct {
	FileContents string
	filePath     string
	done         chan bool
	waitGroup    sync.WaitGroup
	config       config.Config
}

func InitFileWatcher(filePath string, config config.Config) fileWatcher {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("File could not be read: %s", filePath)
	}
	return fileWatcher{FileContents: string(contents), filePath: filePath, config: config, done: make(chan bool)}
}

func (self *fileWatcher) run() {
	ticker := time.NewTicker(3 * time.Second)
	done := make(chan bool)
	self.waitGroup.Add(2)

	go func() {
		for {
			select {
			case <-done:
				self.waitGroup.Done()
				return
			case <-ticker.C:
				self.checkFile()
			}
		}
	}()

	<-self.done
	ticker.Stop()
	done <- true
	self.checkFile() // uno mas because there might be changes
	self.waitGroup.Done()
}

func (self *fileWatcher) checkFile() {
	bytes, err := os.ReadFile(self.filePath)
	if err != nil {
		log.Println("Error reading file - checkFile", self.filePath)
	}

	contents := string(bytes)
	if contents != self.FileContents {
		self.FileContents = contents
		commitFile(self.filePath, self.config)
	}
}

func (self *fileWatcher) stop() {
	self.done <- true
	self.waitGroup.Wait()
}
