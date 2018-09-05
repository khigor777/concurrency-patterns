package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		mu.Lock()
		go count()

		mu.Lock()
		go hello()

	}
	wg.Wait()

}

func count() {
	counter++
	mu.Unlock()
	wg.Done()
}

func hello() {
	fmt.Println(counter)
	mu.Unlock()
	wg.Done()
}
