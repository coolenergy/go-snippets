package io

import (
	"os"
	"text/tabwriter"
)

func tabWriterFunc() {
	writer := tabwriter.NewWriter(os.Stdout, 30, 0, 1, ' ', tabwriter.AlignRight)

	_, err := writer.Write([]byte("First Line\n"))
	if err != nil {
		panic(err.Error())
	}
	_, err = writer.Write([]byte("Second Line\n"))
	if err != nil {
		panic(err.Error())
	}
	err = writer.Flush()
	if err != nil {
		panic(err.Error())
	}
}
