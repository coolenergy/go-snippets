package system

import (
	"fmt"
	"time"
)

func timeDurationDeltaSleep() {
	rateLimit := time.Second * 5
	now := time.Now()
	time.Sleep(time.Second)
	last := time.Now()
	delta := last.Sub(now)

	fmt.Printf("Sleep %f seconds", delta.Seconds())

	if delta < rateLimit {
		sleep := rateLimit - delta
		fmt.Printf("Gonna sleep for %f seconds\n",
			sleep.Seconds())
		time.Sleep(sleep)
		fmt.Printf("Done")
	}
}
