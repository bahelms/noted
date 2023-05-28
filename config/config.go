package config

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// Config holds runtime information
type Config struct {
	LocalStorageDir string
	Editor          string
	AwsProfile      string
	S3BucketName    string
}

// LocalStorage returns the absolute path of the directory used to store files locally
func (cfg *Config) LocalStorage() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, cfg.LocalStorageDir)
}

// LocalFilePath returns the absolute path of the file in local storage
func (cfg *Config) LocalFilePath(filename string) string {
	return filepath.Join(cfg.LocalStorage(), filename)
}
