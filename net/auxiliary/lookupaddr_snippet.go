package auxiliary

import (
	"fmt"
	"net"
)

func LookupAddressSnippet() {
	addresses, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, address := range addresses {
		fmt.Println(address)
	}
}
