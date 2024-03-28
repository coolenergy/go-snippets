package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter = uint64(0)

func produce(in chan string, waitGroup *sync.WaitGroup) {
	value := atomic.AddUint64(&counter, 1)
	producedPassword := fmt.Sprintf("someInput-%d", value)
	in <- producedPassword
	waitGroup.Done()
}

func consume(in chan string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for inputString := range in { // Input input
		if inputString == "done" {
			close(in)
			return
		}
		fmt.Println(inputString)
	}
}

func producerConsumerSnippet() {

	in := make(chan string)
	producerWaitGroup := &sync.WaitGroup{}
	consumerWaitGroup := &sync.WaitGroup{}

	numCpu := runtime.NumCPU()
	for i := 0; i < numCpu; i++ {
		producerWaitGroup.Add(1)
		go produce(in, producerWaitGroup)
	}

	consumerWaitGroup.Add(1)
	go consume(in, consumerWaitGroup)

	producerWaitGroup.Wait()
	in <- "done"
	consumerWaitGroup.Wait()

}
