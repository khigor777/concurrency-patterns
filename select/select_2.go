package main

import "fmt"

const num = 10

func main() {

	ch1 := make(chan int, num)

	for i := 0; i < num; i++ {
		ch1 <- i
	}

	ch2 := make(chan int, 1)
	ch2 <- 3
LOOP:
	for {
		select {
		case val := <-ch1:
			fmt.Println("read from ch 1 ", val)
		case val2 := <-ch2:
			fmt.Println("read from ch2", val2)
		default:
			break LOOP
		}
	}

}
