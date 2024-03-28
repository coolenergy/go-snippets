package io

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func MultipleLoggerLumberJack() {
	if _, err := os.Stat("test_files/loggers"); os.IsNotExist(err) {
		err = os.MkdirAll("test_files/loggers", 644)
		if err != nil {
			panic(err)
		}
	}

	log1 := &lumberjack.Logger{
		Filename:   "test_files/loggers/lumberjack1.txt",
		MaxSize:    5, // in megabytes
		MaxBackups: 3,
		MaxAge:     1,     // in days
		Compress:   false, // Already false by default
	}

	log2 := &lumberjack.Logger{
		Filename:   "test_files/loggers/lumberjack2.txt",
		MaxSize:    5, // in megabytes
		MaxBackups: 3,
		MaxAge:     1,     // in days
		Compress:   false, // Already false by default
	}

	bytesWritten, err := log1.Write([]byte("First Message for logger 1"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Written %d bytes to log file 1\n", bytesWritten)

	_, err = log2.Write([]byte("First Message for logger 2"))
	if err != nil {
		panic(err)
	}

	err = log1.Rotate()
	if err != nil {
		panic(err)
	}

	_, err = log1.Write([]byte("Second Message for logger 1 on second file"))
	if err != nil {
		panic(err)
	}

	err = log1.Close()
	if err != nil {
		panic(err)
	}
	err = log2.Close()
	if err != nil {
		panic(err)
	}
}
