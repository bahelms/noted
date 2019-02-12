package core_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bahelms/noted/config"
	"github.com/bahelms/noted/core"
)

var cfg = config.Config{
	LocalStorageDir: ".noted_tests",
	Editor:          "cat",
}

var fileCases = []struct {
	input    string
	expected string
	content  string
}{
	{"file", "file.txt", "awesome"},
	{"file.any", "file.any", "radical"},
}

func TestOpenFileCreatesNonExistantFilesLocally(t *testing.T) {
	for _, c := range fileCases {
		fp := config.LocalFilePath(cfg, c.expected)
		os.Remove(fp)

		core.OpenFile(cfg, c.input)
		if _, err := os.Stat(fp); os.IsNotExist(err) {
			t.Errorf("%s was not found.", fp)
		}
	}
}

func TestOpenFileDoesNotCreateFilesIfTheyExist(t *testing.T) {
	for _, c := range fileCases {
		expected := []byte(c.content)
		fp := config.LocalFilePath(cfg, c.expected)
		err := ioutil.WriteFile(fp, expected, 0664)
		if err != nil {
			fmt.Printf("WriteFile error: %s -- %v", fp, err)
		}

		core.OpenFile(cfg, c.input)
		actual, _ := ioutil.ReadFile(fp)
		if !bytes.Equal(actual, expected) {
			t.Errorf("Actual: \"%s\"\tExpected: \"%s\"\tFile: %s", actual, expected, c.input)
		}
	}
}

func TestDeleteFileRemovesLocallyStoredFile(t *testing.T) {
	for _, c := range fileCases {
		fp := config.LocalFilePath(cfg, c.expected)
		_, err := os.Create(fp)
		if err != nil {
			fmt.Printf("Create error: %s -- %v", fp, err)
		}

		core.DeleteFile(cfg, c.input)
		if _, err := os.Stat(fp); os.IsExist(err) {
			t.Errorf("%s was not deleted.", fp)
		}
	}
}

func TestDeleteFileHandlesNonExistentFile(t *testing.T) {
	core.DeleteFile(cfg, "unknown")
}
