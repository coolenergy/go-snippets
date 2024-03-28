package text

import (
	"bytes"
	"fmt"
)

func bytesBufferSnippet() {
	var buffer bytes.Buffer

	buffer.WriteString("First Line\n")
	buffer.Write([]byte("Second Line\n"))

	fmt.Println(buffer.String())
}
