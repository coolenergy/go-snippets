package auxiliary

import (
	"context"
	"fmt"
	"net"
	"strings"
)

func DnsLookupCustomResolver() {
	// https://stackoverflow.com/questions/30043248/why-golang-lookup-function-cant-provide-a-server-parameter
	resolvers := []string{
		"1.1.1.1",
		"8.8.8.8",
	}

	i := -1
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (conn net.Conn, err error) {
			i = (i + 1) % len(resolvers)
			d := net.Dialer{}

			var ip net.IP
			// 127.0.0.53:53
			if strings.HasSuffix(address, ":53") {

				ipOrHostnameStr := strings.Split(address, ":")[0]
				ip = net.ParseIP(ipOrHostnameStr)
				if ip != nil {
					conn, err = d.DialContext(ctx, "udp", ipOrHostnameStr)
					return conn, err
				}

			}

			ip = net.ParseIP(address)

			if ip != nil {
				fmt.Printf("Given address is already an ip %s\n", address)

				conn, err = d.DialContext(ctx, "udp", net.JoinHostPort(address, "53"))

				return conn, err
			}

			ipInterface := ctx.Value("ip")
			if ipInterface != nil {
				switch ipInterface.(type) {
				case string:
					fmt.Printf("Using %s as resolver\n", ipInterface)
					conn, err = d.DialContext(ctx, "udp", net.JoinHostPort(ipInterface.(string), "53"))
					return conn, err
				}
			}

			randomResolver := resolvers[i]
			fmt.Printf("Using random resolver %s\n", randomResolver)
			result, err := d.DialContext(ctx, "udp", net.JoinHostPort(randomResolver, "53"))
			return result, err

		},
	}

	//ctx:=context.Background()
	ctx := context.WithValue(context.Background(), "ip", "192.11.111.11")
	ips, _ := resolver.LookupHost(ctx, "starbucks.com")
	for _, ip := range ips {
		fmt.Println(ip)
	}

	ips, _ = resolver.LookupHost(context.TODO(), "starbucks.com")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
