package system

import (
	"fmt"
	"runtime"
)

func getCpuCoreCount() {
	cores := runtime.NumCPU()

	fmt.Printf("This machine has %d CPU cores. \n", cores)

	// maximize CPU usage for maximum performance
	fmt.Printf("maximize CPU usage for maximum performance. %d \n", runtime.GOMAXPROCS(cores))

}
