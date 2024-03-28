package text

import "strings"

func splitSnippet() {
	sentence := "The quick brown fox jumps over the lazy dog"
	parts := strings.Split(sentence, " ")
	for index := range parts {
		println(parts[index])
	}
}
