package io

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func jsonFileDecode() {

	f, err := os.Open("./test_files/users.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)

	users := struct {
		Users []struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			Age    int    `json:"Age"`
			Social struct {
				Facebook string `json:"facebook"`
				Twitter  string `json:"twitter"`
			} `json:"social"`
		} `json:"users"`
	}{}

	err = decoder.Decode(&users)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
