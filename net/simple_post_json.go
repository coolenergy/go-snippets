package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func simpleJsonPost() {

	message := map[string]interface{}{
		"string":  "simple",
		"integer": 10,
		"object": map[string]string{
			"inner_key": "inner_value",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
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
