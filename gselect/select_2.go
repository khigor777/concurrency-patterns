package main

import "fmt"

func main() {
	done := make(chan struct{})
	ch := make(chan int)

	go func(done chan struct{}, ch chan int) {
		value := 0
		for {
			select {
			case ch <- value:
				value++
			case <-done:
				return
			}
		}
	}(done, ch)

	for curVal := range ch {
		fmt.Println(curVal)
		if curVal == 50 {
			done <- struct{}{}
			break
		}
	}
}
