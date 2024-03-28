package text

import (
	"fmt"
	"strings"
)

func replacerSnippet() {
	sentence := "The quick brown fox jumps over the lazy dog"
	replacer := strings.NewReplacer("fox", "dog", "dog", "cat")
	result := replacer.Replace(sentence)
	fmt.Println(result)
}
