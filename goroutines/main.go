package main

import (
	"fmt"
	"time"
)

func printMessage(msg string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(400 * time.Microsecond)
	}
	// when thread complete assign a value to channel
	channel <- "Done!"
}
func main() {
	// goroutines
	// Decalare a channel
	var channel chan string

	go printMessage("Hello World, Geatenious!", channel)
	// wait for channel response
	response := <-channel

	fmt.Println(response)

	// Buffer channel

	logs := make(chan string, 2)

	logs <- "Log 1"
	logs <- "Log 2"
	fmt.Printf("Buffered channel: %v\n", logs)
}
