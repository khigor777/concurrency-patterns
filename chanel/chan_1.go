package main

import "fmt"

func main() {
	ch := make(chan int) //make(chan int, 1) avoid deadlock.

	go func(in <-chan int) {
		v := <-in
		fmt.Println(v)
		fmt.Println("read from chanel")
	}(ch)

	ch <- 10
	//ch <- 100 //if we write second int item, we'll get deadlock. Make your chan with buffer
	fmt.Println("Main goroutine")
}
