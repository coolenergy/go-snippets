package io

import (
	"bufio"
	"fmt"
	"os"
)

func readLinesFromFile() {
	filePath := "./test_files/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		// Notice Println, line does not end with \n
		fmt.Println(line)
	}

}
