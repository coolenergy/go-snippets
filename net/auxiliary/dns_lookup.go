package auxiliary

import (
	"fmt"
	"net"
)

func DnsLookup() {

	addresses, err := net.LookupHost("google.com")
	if err != nil {
		panic(err.Error())
	}
	println("google.com Addresses count: ", len(addresses))
	for _, address := range addresses {
		fmt.Println(address)
	}

	ips, err := net.LookupIP("google.com")
	if err != nil {
		panic(err)
	}
	println("google.com Ips count: ", len(ips))
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
