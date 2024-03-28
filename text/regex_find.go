package text

import (
	"fmt"
	"regexp"
)

func regexFind() {

	const htmlString = `<html>
<head></head>
<body>
<div class="container">
<h2>Go Strings</h2>
......
<h2>Go IO</h2>
......
</div>
</body></html>`

	h2Regex := regexp.MustCompile("<h2>.*</h2>")
	first := h2Regex.FindString(htmlString)

	fmt.Println("First occurrence", first)

	all := h2Regex.FindAllString(htmlString, -1)
	for _, occurrence := range all {
		fmt.Println(occurrence)
	}
}
