package text

import (
	"fmt"
	"regexp"
)

func findAllIndexSnippet() {

	text := `There are many http URLs
1: https://google.com
2: https://twitter.com`

	r := regexp.MustCompile("https?")
	indexes := r.FindAllIndex([]byte(text), -1)

	for i := 0; i < len(indexes); i++ {
		index := indexes[i]
		start := index[0]
		end := index[1]
		part := text[start:end]

		fmt.Printf("text[%d:%d] = %s\n",
			start, end,
			part)
	}
}
