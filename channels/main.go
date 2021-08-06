package main

import (
	"fmt"
	"time"
)

func main() {
	// create a buffered channel
	// with a capacity of 2.
	ch := make(chan string, 2)
	ch <- "water"

	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Resolve goroutine")
		ch <- "anoti"
		fmt.Println(<-ch)
	}()

	ch <- "jose"
	fmt.Println("First")
	fmt.Println(<-ch)
	fmt.Println(len(ch), " ", cap(ch))
	fmt.Println(<-ch, len(ch))
}

//Non initialized chanels must run in a goroutine,
//if its is used in the main thread, an error will ocurr
// func main() {
// 	// create a buffered channel
// 	// with a capacity of 2.
// 	ch := make(chan string, 1)
// 	ch <- "water"

// 	fmt.Println(<-ch)
// 	fmt.Println(len(ch), " ", cap(ch))
// 	fmt.Println(<-ch, len(ch))
// }

//Non initialized chanels must run in a goroutine,
//if that is used in the main thread, an error will ocurr
// func main() {
// 	// create a buffered channel
// 	// with a capacity of 2.
// 	go func() {
// 		fmt.Println("Inside goroutine")
// 		//If doesnt initialize, the goroutine is always stuck
// 		// ch := make(chan string)
// 		ch := make(chan string, 10)
// 		ch <- "water"
// 		fmt.Println("Alread pass value to chanel")
// 		fmt.Println(<-ch)
// 		fmt.Println(len(ch), " ", cap(ch))
// 		fmt.Println(<-ch, len(ch))
// 	}()

// 	time.Sleep(time.Second * 5)

// }
