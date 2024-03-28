package concurrency

import (
	"fmt"
	"sync"
)

var waitGroup = &sync.WaitGroup{}
var number = 0

func inc() {
	number++
	waitGroup.Done()
}

func RaceOnIncrement() {
	// Hope this snippet helps clarify something I was always asking myself in the past.
	// Do we really need to synchronize code that only writes(adds or decrements) but does not read?
	// Yes, because not only read AND write operations can interfere each other, but also:
	// Write AND write can interfere, while you are writing a value, another thread may come to write
	// a different value, result? unexpected behaviour
	for i := 0; i < 600; i++ {
		waitGroup.Add(1)
		go inc()
	}

	waitGroup.Wait()
	fmt.Printf("The expected value was 600, what we got instead is: %d, rerun the demo "+
		"and you will see another value\n",
		number)
}
