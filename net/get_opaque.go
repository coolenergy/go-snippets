package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWithOpaque() {
	// The endpoint is generated with https://requestbin.com/r/eneszlp5orxv6
	request, err := http.NewRequest("GET", "https://eneszlp5orxv6.x.pipedream.net", nil)

	if err != nil {
		panic(err)
	}

	request.URL.Opaque = "/path?param1=value1&param2=value2"
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	bodyString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %s\n %s\n", response.StatusCode, response.Status, bodyString)
}
