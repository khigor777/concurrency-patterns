package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-timer.C:
			fmt.Println("timer.C timeout happend")
		case <-time.After(time.Minute):
			fmt.Println("time after happend")
		case val := <-longQuery():
			if !timer.Stop() {
				<-timer.C
			}
			fmt.Println(val)
		}
	}

}

func longQuery() chan string {
	ch := make(chan string)
	ch <- "test sql"
	return ch
}
