package net

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWithEnvProxy() {
	os.Setenv("HTTP_PROXY", "http://localhost:8080")
	// Disable SSL check
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("https://httpbin.org/get")
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
