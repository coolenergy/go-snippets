package filesystem

import (
	"fmt"
	"os"
)

func fileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileExistsSnippet() {
	filepath := "test_files/test.txt"
	fmt.Printf("Does %s exists? %t", filepath, fileExists(filepath))
}
