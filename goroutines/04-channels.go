package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)

	go func() {
		fmt.Println("from goroutine 1: reading from channel")
		msg := <-messages
		fmt.Println(msg)
		messages <- "done"
	}()

	go func() {
		fmt.Println("from goroutine 2: reading from channel")
		msg := <-messages
		fmt.Println(msg)
	}()

	fmt.Println("from main : writing 1st message to channel")
	messages <- "Hello"
	fmt.Println("from main : writing 2nd message to channel")
	messages <- "World"
	msg := <-messages
	fmt.Println(msg)
	//runtime.Gosched()
}
