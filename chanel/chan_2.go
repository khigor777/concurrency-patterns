package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan<- int) {

		for i := 0; i < 5; i++ {
			fmt.Println("before ", i)
			ch <- i
			fmt.Println("after", i)
		}
		close(ch) //we must close chanel, if we want read it with range
		fmt.Println("finsh write to chanel")

	}(ch)

	for i := range ch {
		fmt.Println("read chanel item: ", i)
	}
}
