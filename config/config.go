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

func New() Config {
	return Config{LocalStorageDir: ".noted", Editor: "nvim", AwsProfile: "noted", S3BucketName: "noted-file-storage"}
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
