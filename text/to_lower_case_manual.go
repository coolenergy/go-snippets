package text

import "fmt"

func ToLowerCaseManual() {
	message := "HELLO WORLD"
	for i := 0; i < len(message); i++ {
		aLower := message[i] | 32
		fmt.Printf("%c", aLower)
	}

	fmt.Println()
}
