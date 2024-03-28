package text

import (
	"fmt"
	"strings"
)

func EqualFold() {
	// A good article to understand this is https://www.digitalocean.com/community/questions/how-to-efficiently-compare-strings-in-go
	// Create 2 equal strings in different ways.
	message := "Hello World"
	lowerString := "hello world"

	// Test the strings for equality.
	if message == lowerString {
		fmt.Println("message == lowerString match")
	}

	if message != lowerString {
		fmt.Println("message != lowerString match")
	}

	if strings.ToLower(message) == lowerString {
		fmt.Println("lower(message) == lowerString match")
	}

	if strings.EqualFold(message, lowerString) {
		fmt.Println("EqualFold match")
	}

}
