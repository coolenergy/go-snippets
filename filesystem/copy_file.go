package filesystem

import (
	"fmt"
	"io"
	"os"
)

func copyFileDemo() {

	sourceFile := "test_files/test.txt"
	destinationFile := "test_files/test2.txt"

	sourceFileStat, err := os.Stat(sourceFile)
	if err != nil {
		panic(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		panic(fmt.Errorf("%s is not a regular file", sourceFile))
	}

	source, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		panic(err)
	}

	defer destination.Close()
	bytesCopied, err := io.Copy(destination, source)

	if err != nil {
		fmt.Printf("Something went wrong %q\n", err)
	} else {
		fmt.Printf("%d bytes copied\n", bytesCopied)
	}
}
