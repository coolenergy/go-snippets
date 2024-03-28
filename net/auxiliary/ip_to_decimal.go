package auxiliary

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func IP2Integer() (int64, error) {
	ip := net.ParseIP("129.168.1.1")

	ip4 := ip.To4()
	if ip4 == nil {
		return 0, fmt.Errorf("illegal: %v", ip)
	}

	bin := make([]string, len(ip4))
	for i, v := range ip4 {
		bin[i] = fmt.Sprintf("%08b", v)
	}
	return strconv.ParseInt(strings.Join(bin, ""), 2, 64)
}
