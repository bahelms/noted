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
	AwsProfile:      "test",
	S3BucketName:    "test-bucket",
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

func createLocalFile(filename string, contents string) string {
	path := cfg.LocalFilePath(filename)
	err := os.WriteFile(path, []byte(contents), 0644)
	if err != nil {
		fmt.Printf("Create error: %s -- %v", path, err)
	}
	return path
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
	// Mock the DeleteExternalFile function to succeed
	originalDeleteExternalFile := core.DeleteExternalFile
	defer func() { core.DeleteExternalFile = originalDeleteExternalFile }()
	core.DeleteExternalFile = func(cfg config.Config, filename string) error {
		return nil
	}

	for _, testCase := range fileCases {
		path := createLocalFile(testCase.expected, "")

		err := core.DeleteFile(cfg, testCase.input)
		if err != nil {
			t.Errorf("Unexpected error deleting file: %v", err)
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("%s was not deleted.", path)
		}
	}
}

func TestDeleteFileReturnsErrorWhenLocalDeletionFails(t *testing.T) {
	// Mock the DeleteExternalFile function to succeed
	originalDeleteExternalFile := core.DeleteExternalFile
	defer func() { core.DeleteExternalFile = originalDeleteExternalFile }()
	core.DeleteExternalFile = func(cfg config.Config, filename string) error {
		return nil
	}

	// Try to delete a non-existent file
	nonExistentFile := "non_existent_file.txt"
	err := core.DeleteFile(cfg, nonExistentFile)
	if err == nil {
		t.Error("Expected error when deleting non-existent file")
	} else if !strings.Contains(err.Error(), "error deleting local file") {
		t.Errorf("Expected error about local file deletion, got: %v", err)
	}
}

func TestDeleteFileReturnsErrorWhenRemoteDeletionFails(t *testing.T) {
	// Create a file that will be deleted locally
	path := createLocalFile("test.txt", "")
	defer os.Remove(path)

	// Mock the DeleteExternalFile function to return an error
	originalDeleteExternalFile := core.DeleteExternalFile
	defer func() { core.DeleteExternalFile = originalDeleteExternalFile }()
	core.DeleteExternalFile = func(cfg config.Config, filename string) error {
		return fmt.Errorf("mock S3 deletion error")
	}

	err := core.DeleteFile(cfg, "test.txt")
	if err == nil {
		t.Error("Expected error when remote deletion fails")
	} else if !strings.Contains(err.Error(), "error deleting remote file") {
		t.Errorf("Expected error about remote file deletion, got: %v", err)
	}

	// Verify local file was still deleted
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Error("Local file should have been deleted despite remote deletion failure")
	}
}

func TestListFilesPrintsAllLocalFilesToStdout(t *testing.T) {
	var expectedFiles [2]string
	for i, testCase := range listFilesCases {
		createLocalFile(testCase.expected, "")
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
