package text

import (
	"fmt"
	"strings"
)

func strReplace() {
	sentence := "the quick brown fox jumps over the lazy dog"
	result := strings.Replace(sentence, "the", "some", -1)
	fmt.Println(result)

	result = strings.Replace(sentence, "the", "some", 1)
	fmt.Println(result)
}
