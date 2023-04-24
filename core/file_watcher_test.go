package core_test

import (
	"testing"

	"github.com/bahelms/noted/core"
)

func TestInitFileWatcherReadsFileContents(t *testing.T) {
	path := createLocalFile("fileWatcherTest.txt", "hey there")

	watcher := core.InitFileWatcher(path)
	if watcher.FileContents != "hey there" {
		t.Errorf("File contents not read correctly")
	}
}

func TestRunSavesFileToExternalStoreWhenDiffOccurs(t *testing.T) {
}

func TestRunDoesNotSaveFileWithoutDiff(t *testing.T) {
}
