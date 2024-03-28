package process

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"time"
)

func procStdInStdOut() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "cmd.exe"
	} else {
		cmd = "/bin/bash"
	}
	proc := exec.Command(cmd)
	stdin, err := proc.StdinPipe()
	if err != nil {
		panic(err.Error())
	}

	defer stdin.Close()

	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = proc.Start()
	if err != nil {
		panic(err.Error())
	}

	_, err = io.WriteString(stdin, "whoami\n")
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(time.Second * 2)
	proc.Process.Kill()
}
