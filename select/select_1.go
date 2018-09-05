package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)
	ch1 <- 1

	select {
	case val := <-ch1:
		fmt.Println(val)
	case ch2 <- 1:
		fmt.Println("put 1 to ch2")
	default:
		fmt.Println("default case")
	}

}
