package process

import (
	"fmt"
	"os"
)

func getSelfPid() {
	pid := os.Getpid()
	fmt.Printf("Self process id: %d\n", pid)
}
