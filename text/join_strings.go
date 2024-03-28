package text

import (
	"strings"
)

func joinStrings() {
	var lines = []string{
		"First Line",
		"Second Line",
		"Third Line",
	}
	text := strings.Join(lines, "\n")
	println(text)

}
