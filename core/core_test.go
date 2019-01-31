package core

import "testing"

func TestOpenFile(t *testing.T) {
	result := OpenFile("some_file")
	if result != "some_file" {
		t.Error("Test failed")
	}
}
