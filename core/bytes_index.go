package core

import (
	"bytes"
	"fmt"
)

func BytesIndex() {
	sentence := []byte("We are learning how to use bytes.Index function with bytes")

	indexOfLearning := bytes.Index(sentence, []byte("learning"))
	if indexOfLearning == -1 {
		fmt.Printf(`"learning" word is at index %d`, indexOfLearning)
	}

	indexOfStrings := bytes.Index(sentence, []byte("strings"))
	if indexOfStrings == -1 {
		fmt.Printf(`"strings" word was not found %d`, indexOfStrings)
	}
}
