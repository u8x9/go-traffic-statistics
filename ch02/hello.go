package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go func() {
		msg <- "hello goroutine!"
	}()
	go func() {
		time.Sleep(time.Second)
		str := <-msg
		str += "more and more"
		msg <- str
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(<-msg)
	fmt.Println("hello world")
}
