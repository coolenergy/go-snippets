package text

import (
	"fmt"
	"regexp"
)

func regexReplace() {
	lines := "1. First item\n2. Second Item\n3. Third Item\n"

	regex := regexp.MustCompile("[0-9]+\\.")
	out := regex.ReplaceAllString(lines, "-")
	fmt.Println(out)

}
