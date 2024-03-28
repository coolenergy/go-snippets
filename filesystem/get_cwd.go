package filesystem

import (
	"fmt"
	"log"
	"os"
)

func GetCwdSnippet() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
