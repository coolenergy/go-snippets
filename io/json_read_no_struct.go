package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func jsonReadNonStruct() {

	// Open our jsonFile
	jsonFile, err := os.Open("test_files/json/users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string][]map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result["users"])

	fmt.Println(result["users"][0]["name"])

}
