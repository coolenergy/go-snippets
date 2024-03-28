package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/**
This snippet shows how to use condition variables.
Condition variables are a very well known concurrency primitives used to synchronize threads
when some kind of signaling is needed.
In these scenarios we use together two primitives:
Condition Variable and a Mutex
Since they go together, you should consider wrapping them in a single struct
And creating wrapper functions that know how to use them both, if a function only deals with one of them
forgetting the other, you may end up in a deadlock situation.
*/
type Record struct {
	sync.Mutex
	data string

	cond *sync.Cond
}

func DataSource() *Record {
	r := Record{}
	r.cond = sync.NewCond(&r)
	return &r
}

func ConditionVariableUsage() {
	var wg sync.WaitGroup

	rec := DataSource()
	wg.Add(1)
	go func(rec *Record) {
		defer wg.Done()
		fmt.Println("Goroutine tries to lock the mutex")
		rec.Lock()
		fmt.Println("Goroutine locks the mutex")
		fmt.Println("Goroutine waits a signal, so mutex is released")
		rec.cond.Wait()
		// Once we are waked up from wait, we own the lock, don't forget to release it
		fmt.Println("Goroutine waken up, will release the mutex")
		rec.Unlock()
		fmt.Println("Goroutine released the mutex")
		fmt.Println("Data: ", rec.data)
		return
	}(rec)

	time.Sleep(2 * time.Second)
	fmt.Println("Main thread will acquire the mutex")
	rec.Lock()
	fmt.Println("Main thread acquired the mutex")
	rec.data = "gopher"
	fmt.Println("Main thread will release the mutex")
	rec.Unlock()
	fmt.Println("Main thread released the mutex and will signal")

	rec.cond.Signal()
	fmt.Println("Main thread signaled")
	wg.Wait() // wait till all goroutine completes
}
