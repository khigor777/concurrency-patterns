package main

import (
	"fmt"
	"time"
)

func main() {
	resCh := getComments()
	time.Sleep(time.Second)
	fmt.Println("Do some work")
	res := <-resCh
	fmt.Println(res)
}

func getComments() chan string {
	ch := make(chan string, 1) // must be buffered
	go func(out chan string) {
		time.Sleep(time.Second * 2)
		fmt.Println("func ready return result")
		ch <- "func done work"
	}(ch)
	return ch
}
