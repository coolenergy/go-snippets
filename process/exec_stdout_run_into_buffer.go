package process

import (
	"bytes"
	"fmt"
	"os/exec"
)

func execStdOutRunIntoBuffer() {
	process := exec.Command("ls", "-l")
	stdOut := bytes.NewBuffer([]byte{})
	process.Stdout = stdOut
	err := process.Run() // Run and wait to complete, .Start() would not wait to be completed
	if err != nil {
		fmt.Println(err)
	}

	if process.ProcessState.Success() {
		fmt.Println(stdOut.String())
	}
}
