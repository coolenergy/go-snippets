package process

import (
	"fmt"
	"os/exec"
)

func execCommand() {
	prc := exec.Command("ls", "-p")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
