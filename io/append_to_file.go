package io

import (
	"log"
	"os"
)

func appendToFile() {

	file, err := os.OpenFile("test_files/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("text to append\n"); err != nil {
		log.Println(err)
	}
}
