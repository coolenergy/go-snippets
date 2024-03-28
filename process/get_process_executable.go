package process

import (
	"fmt"
	"os"
	"path/filepath"
)

func getProcessExecPath() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Application path %s\n", exe)

	exePath := filepath.Dir(exe)
	fmt.Println("DirectoryPath :" + exePath)
}
