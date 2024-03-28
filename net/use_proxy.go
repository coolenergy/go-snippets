package net

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetWithProxy() {

	proxyUrl, err := url.Parse("http://localhost:8080")
	// // http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport = &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

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
