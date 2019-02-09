package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bahelms/noted/config"
)

// OpenFile opens a file
func OpenFile(cfg config.Config, filename string) string {
	ensureLocalStorage(cfg)
	fp := config.LocalFilePath(cfg, ensureExtension(filename))
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if _, err := os.Create(fp); err != nil {
			log.Fatalf("Failed creating %s -- %s", fp, err)
		}
	}

	// set watcher on filepath
	// open file with nvim
	// cmd := exec.Command("nvim")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal("nvim errored:", err)
	// }
	return filename
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
