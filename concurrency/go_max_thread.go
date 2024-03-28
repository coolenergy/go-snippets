package concurrency

import (
	"fmt"
	"runtime"
)

func goMaxThreads() {
	// retrieves the max number of goroutines that can be run simultaneously, which is the CPU core count
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Number of cores %d\n",
		maxProcs)
}
