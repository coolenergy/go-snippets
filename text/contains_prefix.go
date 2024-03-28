package text

import (
	"fmt"
	"strings"
)

func containsSnippet() {
	const sentence = "The quick brown fox jumps over the lazy dog"
	contain := strings.Contains(sentence, "fox")
	fmt.Printf("Contains fox? %t \n", contain)

	contain = strings.Contains(sentence, "dog")
	fmt.Printf("Contains dog? %t \n", contain)
}
