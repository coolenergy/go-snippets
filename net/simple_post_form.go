package net

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func simplePostForm() {

	formData := url.Values{
		"username": {"MelarDev"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("{")
	for key, value := range result {
		switch value.(type) {
		case string:
			fmt.Printf("\t%s: %s\n", key, value)
		case map[string]interface{}:
			fmt.Printf("\t%s: {\n", key)
			for innerKey, innerValue := range value.(map[string]interface{}) {
				fmt.Printf("\t\t%s: %s\n", innerKey, innerValue)
			}
			fmt.Println("\t}")
		}
	}

	fmt.Println("}")
}
