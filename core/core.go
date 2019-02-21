package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bahelms/noted/config"
)

// OpenFile opens a file
func OpenFile(cfg config.Config, filename string) {
	// setup file
	ensureLocalStorage(cfg)
	fp := config.LocalFilePath(cfg, ensureExtension(filename))
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if _, err := os.Create(fp); err != nil {
			log.Fatalf("Failed creating %s -- %s", fp, err)
		}
	}

	// set watcher on file

	// open text editor
	externalCommand := exec.Command(cfg.Editor, fp)
	externalCommand.Stdin = os.Stdin
	externalCommand.Stdout = os.Stdout
	if err := externalCommand.Run(); err != nil {
		log.Fatalf("%s errored: %s", cfg.Editor, err)
	}
}

// DeleteFile removes the specified file locally
func DeleteFile(cfg config.Config, filename string) {
	fp := config.LocalFilePath(cfg, ensureExtension(filename))
	os.Remove(fp)
}

// ListFiles prints all local files to STDOUT
func ListFiles(cfg config.Config) {
	files, _ := ioutil.ReadDir(config.LocalStorage(cfg))
	for _, file := range files {
		log.Println(file.Name())
	}
}

func ensureLocalStorage(cfg config.Config) {
	localStorage := config.LocalStorage(cfg)
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
