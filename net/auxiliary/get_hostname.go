package auxiliary

import (
	"fmt"
	"os"
)

func GetHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)
}
