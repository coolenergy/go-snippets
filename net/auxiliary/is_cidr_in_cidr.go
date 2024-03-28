package auxiliary

import (
	"fmt"
	"net"
)

func IsCIDRInCIDR() {

	A := "172.17.0.0/16"
	B := "172.17.0.2/16"

	ipA, ipnetA, _ := net.ParseCIDR(A)
	ipB, ipnetB, _ := net.ParseCIDR(B)

	fmt.Println("Network address A: ", A)
	fmt.Println("IP address      B: ", B)
	fmt.Println("ipA              : ", ipA)
	fmt.Println("ipnetA           : ", ipnetA)
	fmt.Println("ipB              : ", ipB)
	fmt.Println("ipnetB           : ", ipnetB)

	fmt.Printf("\nDoes A (%s) contain: B (%s)?\n", ipnetA, ipB)

	if ipnetA.Contains(ipB) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
