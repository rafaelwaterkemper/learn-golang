package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGTERM)

	go func() {
		fmt.Println("goroutine", os.Getpid(), " ", os.Getppid())
		<-channel
		fmt.Println("Receive a sigTerm")
		fmt.Println("Now, sending a kill signal")
		syscall.Kill(os.Getpid(), syscall.SIGKILL)
	}()

	for {
		fmt.Println("printando", os.Getpid(), " ", os.Getppid())
		time.Sleep(time.Second * 5)
	}
}
