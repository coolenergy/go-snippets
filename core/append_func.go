package core

import "fmt"

func AppendToByteArray() {
	templateUsername := []byte("Melar")
	// Method 1
	username := append(templateUsername, ' ')
	username = append(username, 'D')
	username = append(username, 'e')
	username = append(username, 'v')

	fmt.Println(string(username))

	// Method 2
	username2 := append(templateUsername, " Dev"...)

	fmt.Println(string(username2))
}
