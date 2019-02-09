package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bahelms/noted/config"
)

// OpenFile opens a file
func OpenFile(cfg config.Config, filename string) string {
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
	editor := "nvim"
	cmd := exec.Command(editor, fp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s errored: %s", editor, err)
	}
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
