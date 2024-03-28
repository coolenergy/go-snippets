package text

import (
	"regexp"
)

func regexSplit() {
	const sentences = "Username: Melar\nUsername:Dev"
	// change to :\\s+ and you will see a difference
	parts := regexp.MustCompile(":\\s*").Split(sentences, -1)
	for _, part := range parts {
		println(part)
	}

}
