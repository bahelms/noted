package core

import "fmt"

// OpenFile opens a file
func OpenFile(filename string) string {
	fmt.Println("File to open:", filename)
	return filename
}
