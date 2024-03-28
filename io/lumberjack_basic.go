package io

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func LumberJackBasic() {
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

	bytesWritten, err := log1.Write([]byte("First Message for logger 1, very basic"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of bytes written %d\n", bytesWritten)

}
