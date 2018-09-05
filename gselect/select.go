package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2

	ch2 <- 3
	ch2 <- 4

LOOP:
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
		default:
			break LOOP
		}
	}

}
