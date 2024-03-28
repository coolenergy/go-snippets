package net

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getRequest() {

	client := http.Client{}
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

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
