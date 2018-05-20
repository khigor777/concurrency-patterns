package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	input := make(chan string, 2)

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(input)
	}

	items := []string{
		"one", "two", "five", "ten", "one", "two", "five", "ten",
	}
	ticks := time.Tick(time.Second)
	for range ticks {
		for _, v := range items {
			input <- v
		}
	}

	close(input)

}

func worker(ch chan string) {
	for item := range ch {
		fmt.Println(item)
	}
}
