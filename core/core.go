package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bahelms/noted/config"
)

// OpenFile opens a file
func OpenFile(cfg config.Config, filename string) {
	// setup file
	ensureLocalStorage(cfg)
	fp := cfg.LocalFilePath(ensureExtension(filename))
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if _, err := os.Create(fp); err != nil {
			log.Fatalf("Failed creating %s -- %s", fp, err)
		}
	}

	// set watcher on file
	fileWatcher := InitFileWatcher(fp, cfg)
	go fileWatcher.run()
	defer fileWatcher.stop()

	// open text editor
	cmd := exec.Command(cfg.Editor, fp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Panicf("%s Command failed - errored: %s", cfg.Editor, err)
	}
}

// DeleteFile removes the specified file locally
func DeleteFile(cfg config.Config, filename string) {
	fp := cfg.LocalFilePath(ensureExtension(filename))
	os.Remove(fp)
}

// ListFiles prints all local files to STDOUT
func ListFiles(cfg config.Config) {
	files, _ := ioutil.ReadDir(cfg.LocalStorage())
	for _, file := range files {
		log.SetFlags(0)
		ts := file.ModTime().Format("01/02/2006")
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		log.Printf("%s - %s\n", ts, name)
	}
}

func SyncFiles(cfg config.Config) {
	downloadAllFiles(cfg)
}

func ensureLocalStorage(cfg config.Config) {
	localStorage := cfg.LocalStorage()
	if _, err := os.Stat(localStorage); os.IsNotExist(err) {
		os.Mkdir(localStorage, os.ModePerm)
	}
}

func ensureExtension(name string) string {
	if filepath.Ext(name) == "" {
		return fmt.Sprintf("%s.txt", name)
	}
	return name
}
