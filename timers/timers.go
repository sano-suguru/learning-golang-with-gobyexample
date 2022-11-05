package main

import (
	"fmt"
	"time"
)

func main() {
	timers1 := time.NewTimer(2 * time.Second)

	<-timers1.C
	fmt.Println("Timer 1 fired")

	timers2 := time.NewTimer(time.Second)
	go func() {
		<-timers2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timers2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
