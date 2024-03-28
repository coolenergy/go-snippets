package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFormFile() {

	// Open the file
	file, err := os.Open("test_files/test.txt")

	if err != nil {
		log.Fatalln(err)
	}
	// Close the file later
	defer file.Close()

	// Buffer to store our request body as bytes
	var requestBody bytes.Buffer

	// Create a multipart writer
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Initialize the file field
	fileWriter, err := multiPartWriter.CreateFormFile("file_to_upload", "my_file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Copy the actual file content to the field field's writer
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err)
	}

	// Populate other fields
	fieldWriter, err := multiPartWriter.CreateFormField("username")
	if err != nil {
		panic(err)
	}

	_, err = fieldWriter.Write([]byte("MelarDev"))
	if err != nil {
		panic(err)
	}

	// We completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	multiPartWriter.Close()

	// By now our original request body should have been populated, so let's just use it with our custom request
	req, err := http.NewRequest("POST", "https://httpbin.org/post", &requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	// We need to set the content type from the writer, it includes necessary boundary as well
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

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
