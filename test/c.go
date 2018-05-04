package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	timer := time.NewTimer(time.Second * 1)
	aaa := make(chan struct{})
	go func() {

		for {
			timer.Reset(1* time.Second)
			select {
			case <-aaa:
				//timer.Stop()
				fmt.Println("a")
				return

			case <-timer.C:
				fmt.Println("haha")
				//default:
				//	fmt.Println("a")
			}

		}
	}()

	fmt.Println(runtime.NumGoroutine())
	time.Sleep(3*time.Second)
	close(aaa)
	time.Sleep(1*time.Second)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(30*time.Second)
}
