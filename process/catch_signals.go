package process

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func catchSignals() {
	// Create a channel
	signalChannel := make(chan os.Signal, 1)

	// Bind our signalChannel to signals
	signal.Notify(signalChannel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	// Create another channel
	exitChannel := make(chan int)
	go func() {
		// Wait until we get something from the signalChannel
		// which will be one of our signal
		currentSignal := <-signalChannel
		switch currentSignal {

		// We have some signal, feed an exit code into the exitChannel
		case syscall.SIGINT:
			fmt.Println("SIGINT signal received! CTRL + C entered")
			exitChannel <- 1

		case syscall.SIGTERM:
			fmt.Println("SIGTERM signal received")
			exitChannel <- 1

		case syscall.SIGKILL:
			fmt.Println("SIGKILL received")
			exitChannel <- 1

		case syscall.SIGQUIT:
			fmt.Println("SIGQUIT received")
			exitChannel <- 1
		}
	}()

	// Wait until we get something into the exitChannel
	code := <-exitChannel
	fmt.Println("Error code", code)
}
