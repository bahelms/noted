package core_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

var listFilesCases = []struct {
	input    string
	expected string
	content  string
}{
	{"file", "file", "awesome"},
	{"file.any", "file", "radical"},
}

func createLocalFile(filename string) string {
	fp := cfg.LocalFilePath(filename)
	_, err := os.Create(fp)
	if err != nil {
		fmt.Printf("Create error: %s -- %v", fp, err)
	}
	return fp
}

func captureOutput(fn func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	fn()
	log.SetOutput(os.Stdout)
	return buf.String()
}

func TestOpenFileCreatesNonExistantFilesLocally(t *testing.T) {
	for _, c := range fileCases {
		fp := cfg.LocalFilePath(c.expected)
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
		fp := cfg.LocalFilePath(c.expected)
		err := ioutil.WriteFile(fp, expected, 0664)
		if err != nil {
			t.Errorf("WriteFile error: %s -- %v", fp, err)
		}

		core.OpenFile(cfg, c.input)
		actual, _ := ioutil.ReadFile(fp)
		if !bytes.Equal(actual, expected) {
			t.Errorf("Actual: \"%s\"\tExpected: \"%s\"\tFile: %s", actual, expected, c.input)
		}
	}
}

func TestDeleteFileRemovesLocallyStoredFile(t *testing.T) {
	for _, testCase := range fileCases {
		path := createLocalFile(testCase.expected)

		core.DeleteFile(cfg, testCase.input)
		if _, err := os.Stat(path); os.IsExist(err) {
			t.Errorf("%s was not deleted.", path)
		}
	}
}

func TestListFilesPrintsAllLocalFilesToStdout(t *testing.T) {
	var expectedFiles [2]string
	for i, testCase := range listFilesCases {
		createLocalFile(testCase.expected)
		expectedFiles[i] = testCase.expected
	}

	output := captureOutput(func() {
		core.ListFiles(cfg)
	})

	for _, expected := range expectedFiles {
		if !strings.Contains(output, expected) {
			t.Errorf("Actual %s -- expected %s", output, expected)
		}
	}
}
