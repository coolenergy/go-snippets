package process

import (
	"bufio"
	"fmt"
	"os/exec"
)

func procStdOutScanner() {

	proc := exec.Command("ping", "google.com")

	// The process output is obtained
	// in form of io.ReadCloser. The underlying
	// implementation use the os.Pipe
	procStdOut, _ := proc.StdoutPipe()
	defer procStdOut.Close()

	// Start the process
	err := proc.Start()
	if err != nil {
		panic("Error!")
	}

	scanner := bufio.NewScanner(procStdOut)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		println(scanner.Err().Error())
	}
}
