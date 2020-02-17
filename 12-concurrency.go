package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main started")
	go func() {
		fmt.Println("worker executed")
	}()
	fmt.Println("main completed")
	/* fmt.Scanln()
	fmt.Println("done") */
	runtime.Gosched()
}
