package core_test

import (
	"testing"

	"github.com/bahelms/noted/config"
	"github.com/bahelms/noted/core"
)

func TestInitFileWatcherReadsFileContents(t *testing.T) {
	path := createLocalFile("fileWatcherTest.txt", "hey there")

	cfg := config.Config{
		LocalStorageDir: ".noted_tests",
		Editor:          "cat",
	}
	watcher := core.InitFileWatcher(path, cfg)
	if watcher.FileContents != "hey there" {
		t.Errorf("File contents not read correctly")
	}
}
