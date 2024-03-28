package io

import (
	"io"
	"os"
)

func writeToFile() {
	var err error
	var fileWriter *os.File
	fileWriter, err = os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer fileWriter.Close()

	_, err = io.WriteString(fileWriter, "First Line\n")
	_, err = io.WriteString(fileWriter, "Second Line\n")

}
