package io

import (
	"io"
	"log"
	"os"
)

func MultipleLoggerStd() {
	if _, err := os.Stat("test_files/loggers"); os.IsNotExist(err) {
		err = os.MkdirAll("test_files/loggers", 644)
		if err != nil {
			panic(err)
		}
	}

	f1, err := os.Create("test_files/loggers/f1.txt")
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	f2, err := os.Create("test_files/loggers/f2.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	w := io.MultiWriter(os.Stdout, f1, f2)
	loggerF1F2 := log.New(w, "Info - ", log.LstdFlags)

	loggerF1F2.Println("A message for f1 and f2")

	f3, err := os.Create("test_files/loggers/f3.txt")
	if err != nil {
		panic(err)
	}

	defer f3.Close()

	loggerF3 := log.New(f3, "Debug - ", log.LstdFlags)
	loggerF3.Println("A message for f3")

	f4, err := os.Create("test_files/loggers/f4.txt")
	if err != nil {
		panic(err)
	}

	defer f4.Close()
	w2 := io.MultiWriter(f3, f4)

	loggerF4 := log.New(w2, "Warning - ", log.LstdFlags)
	loggerF4.Println("A message for f3 and f4")

}
