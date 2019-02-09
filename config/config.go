package config

import (
	"os/user"
	"path/filepath"
)

// Config holds runtime information
type Config struct {
	LocalStorageDir string
}

// LocalStorage returns the absolute path of the directory used to store files locally
func LocalStorage(cfg Config) string {
	var usr, _ = user.Current()
	return filepath.Join(usr.HomeDir, cfg.LocalStorageDir)
}

// LocalFilePath returns the absolute path of the file in local storage
func LocalFilePath(cfg Config, filename string) string {
	return filepath.Join(LocalStorage(cfg), filename)
}
