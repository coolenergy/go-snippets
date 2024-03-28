package net

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getWithCustomHost() {
	// forward requests to proxy to inspect if this worked
	os.Setenv("HTTP_PROXY", "http://localhost:8080")
	// Disable SSL check
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", "https://34.206.241.1/access/unauthenticated", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = "booking.zendesk.com"
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}

func GetWithCustomHost() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	bytesRepresentation, err := json.Marshal(map[string]interface{}{
		"text": "Test",
	})

	// if err != nil {panic(err)}

	// https://httpbin.org/response-headers
	req, err := http.NewRequest("POST",
		"https://3.220.112.94/response-headers",
		bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		panic(err)
	}

	req.Host = "httpbin.org"
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

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
