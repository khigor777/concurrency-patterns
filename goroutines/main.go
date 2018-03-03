package main

import "fmt"

const (
	goroutineNum = 4
	iterationNum = 4
)

func main() {

	for i := 0; i < goroutineNum; i++ {
		go doSomething(i)
		//runtime.Gosched() Gosched yields the processor, allowing other goroutines to run.
	}
	fmt.Scanln() //wait output
}

func doSomething(goroutine int) {
	for k := 0; k < iterationNum; k++ {
		fmt.Printf("gorutine:%d, currentJob:%d\n", goroutine, k)
	}
}
