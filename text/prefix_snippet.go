package text

import (
	"fmt"
	"strings"
)

func prefixSnippet() {
	sentence := "The quick brown fox jumps over the lazy dog"
	prefix := "The"
	hasPrefix := strings.HasPrefix(sentence, prefix)
	fmt.Printf("Starts with 'The' ? %t\n", hasPrefix)

	suffix := "Dog"
	hasSuffix := strings.HasSuffix(sentence, suffix)
	fmt.Printf("Ends with Dog? %t\n", hasSuffix)
}
