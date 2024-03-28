package process

import (
	"bufio"
	"fmt"
	"os"
)

func readStdInScanner() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Write a message> ")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
