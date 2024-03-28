package process

import (
	"fmt"
	"os/exec"
	"time"
)

func processFinishInfo() {
	cmd := exec.Command("ping", "google.com")
	err := cmd.Run()
	if err != nil {
		panic("Error running app")
	}
	// Any of the following info can only be retrieved after we finished execution
	fmt.Printf("PID: %d\n", cmd.ProcessState.Pid())
	fmt.Printf("Process took: %d ms\n", cmd.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited sucessfuly : %t\n", cmd.ProcessState.Success())
}
