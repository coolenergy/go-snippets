package io

import (
	"bufio"
	"fmt"
	"os"
)

func writeLinesToFile() {
	lines := []string{"First", "Second", "Third"}
	filePath := "test_files/test.txt"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
