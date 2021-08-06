// Timers are for when you want to do something once in the future - tickers are for when
// you want to do something repeatedly at regular intervals. Here’s an example of a ticker
// that ticks periodically until we stop it.

// Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll
// use the select builtin on the channel to await the values as they arrive every 500ms.

package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
