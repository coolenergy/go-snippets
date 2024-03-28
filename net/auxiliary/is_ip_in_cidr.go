package auxiliary

import (
	"fmt"
	"net"
)

func IsIpInCIDR() {
	cidrString := "192.168.1.0/24"
	ipAddress := "192.168.1.130"

	_, cidrNetBlock, _ := net.ParseCIDR(cidrString)
	ipB := net.ParseIP(ipAddress)

	fmt.Printf("\nDoes cidrString (%s) contain: ipAddress (%s)?\n", cidrNetBlock, ipB)

	if cidrNetBlock.Contains(ipB) {
		fmt.Printf("%s Is contained in CIDR %s\n", ipAddress, cidrString)
	} else {
		fmt.Printf("No no, %s Is not contained by CIDR %s\n", ipAddress, cidrString)
	}
}
