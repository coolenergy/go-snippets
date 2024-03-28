package auxiliary

import (
	"fmt"
	"regexp"
)

// https://www.socketloop.com/tutorials/golang-find-ip-address-from-string/?utm_source=socketloop&utm_medium=tutesidebar
func findIP(input string) string {
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}

func GetIpFromString() {

	var ip string

	str := "this is a string that contain IP address 192.168.0.1 and you need to parse it"

	ip = findIP(str)

	fmt.Printf("The IP address in string is %q \n", ip)

}
