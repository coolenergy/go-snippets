package net

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
)

var resolver *net.Resolver

func InitResolver() {
	// https://stackoverflow.com/questions/30043248/why-golang-lookup-function-cant-provide-a-server-parameter
	resolvers := []string{
		"1.1.1.1",
		"8.8.8.8",
	}

	i := -1
	resolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (conn net.Conn, err error) {
			i = (i + 1) % len(resolvers)
			d := net.Dialer{}

			fmt.Printf("DNS lookup: %s\n", address)

			var ip net.IP
			// 127.0.0.53:53
			if strings.HasSuffix(address, ":53") {
				ipOrHostnameStr := strings.Split(address, ":")[0]
				ip = net.ParseIP(ipOrHostnameStr)
				if ip != nil {
					conn, err = d.DialContext(ctx, "udp", address)
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
}

func InitHttpClient() {
	dialer := &net.Dialer{}

	http.DefaultTransport.(*http.Transport).DialContext =
		func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
			// pin the IP (--resolve in cURL) so we don't fuzz our DNS resolver
			addrParts := strings.Split(addr, ":")
			if len(addrParts) == 2 {
				// Use the application resolver
				host := addrParts[0]
				port := addrParts[1]
				addresses, err := resolver.LookupHost(context.TODO(),
					host)

				if err == nil {
					if len(addresses) > 0 {
						addr = net.JoinHostPort(addresses[rand.Intn(len(addresses))], port)
					} else {
						message := fmt.Sprintf("len(addresses) is 0 %s", host)
						fmt.Println(message)
						log.Println(message)
					}
				}

			} else {
				// TODO: Log this unexpected scenario
				message := fmt.Sprintf("Unexpected addrParts len %s\n", addrParts)
				fmt.Println(message)
				log.Println(message)
			}

			conn, err = dialer.DialContext(ctx, network, addr)
			return conn, err
		}
}

func GetWithCustomResolverAndHost() {

	InitResolver()
	InitHttpClient()

	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = "localhost"
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// POST
	bytesRepresentation, err := json.Marshal(map[string]interface{}{
		"text": "Test",
	})

	// if err != nil {panic(err)}

	// https://httpbin.org/headers
	req, err = http.NewRequest("POST",
		"https://httpbin.org/post",
		bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		panic(err)
	}

	req.Host = "localhost"
	req.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyBytes))
	resp.Body.Close()
}
