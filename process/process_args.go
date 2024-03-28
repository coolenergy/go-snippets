package process

import (
	"fmt"
	"os"
)

func processArgsSnippet() {
	processArgs := os.Args

	fmt.Println(processArgs)

	appName := processArgs[0]
	fmt.Printf("Application name: %s \n", appName)

	appArgs := processArgs[1:]
	fmt.Println(appArgs)

	for idx, arg := range appArgs {
		fmt.Printf("Arg %d = %s\n", idx, arg)
	}

}
