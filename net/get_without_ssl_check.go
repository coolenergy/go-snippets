package net

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWithoutSSLCheck() {
	// Create a new transport, inside create a config which ignores SSL check
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transCfg}

	response, err := client.Get("https://httpbin.org/get")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(os.Stdout, string(body))

}
