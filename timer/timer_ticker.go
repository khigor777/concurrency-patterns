package main

import (
	"fmt"
	"time"
)

func main() {
	runAllTimeTicker()
	//tickerNeedStop()
}

func runAllTimeTicker() {
	tick := time.Tick(time.Second)
	for t := range tick {
		fmt.Println(t)
	}
}

func tickerNeedStop() {
	ticker := time.NewTicker(time.Second * 1)
	var i int
	for t := range ticker.C {
		i++
		fmt.Println(t, "tick")
		if i == 5 {
			//need stop or lick memory
			ticker.Stop()
			break
		}
	}
	fmt.Println("total ", i)
}
