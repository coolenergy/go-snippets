package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ioReadChunksSequence() {
	resp, err := http.Get("https://edition.cnn.com/")
	if err != nil {
		panic(err)
	}

	respContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	filename := "test_files/big_file.html"
	err = ioutil.WriteFile(filename, respContent, 0644)
	if err != nil {
		panic(err)
	}

	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fd)
	var bufferString string
	fmt.Printf("Buffer Info: Address %p\n", &bufferString)

	buffer32kb := make([]byte, 32*1024)
	// err == nil if read == len(buffer32kb), that means there may be more data to read
	// err == EOF if read is 0 that means either file is empty if first time, or we read it all if not first
	// err == ErrUnexpectedEOF when read > 0 && read < len(buffer32kb)
	read := 0
	for i := 0; ; i++ {
		read, err = io.ReadFull(reader, buffer32kb)
		fmt.Printf("Iteration %d: %d bytes read; %v\n", i, read, err)
		if err == io.EOF {
			break
		}
	}

	fmt.Println("Bytes read: ", read, "Error: ", err)
	// Notice the pointer of bufferString did not change, so this is a very performant way
	// of reading big chunks, instead of mapping it fully into memory we read chunk to chunk.
	// the downside of course we only have one chunk at a time, the other chunks will be gone
	// and replaced with the new ones as they come.
	fmt.Printf("Buffer Info: Address %p\n", &bufferString)
}
