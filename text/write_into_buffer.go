package text

import (
	"bytes"
	"fmt"
)

func writeIntoBuffer() {
	buffer := bytes.Buffer{}
	buffer.WriteString("Hello Bytes Golang\n")
	buffer.WriteString("Golang is such a fun thing\n")

	fmt.Println(buffer.String())
}
