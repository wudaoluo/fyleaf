package main

import (
	//"sync/atomic"
	"fmt"
	"time"
	"sync"
)


var connCount1 int32
var mu sync.Mutex

func main() {
	//atomic.StoreInt32(&connCount, 0)

	//fmt.Println(atomic.LoadInt32(&connCount))
	n := 0
	go func() {

		for {
			n++
			mu.Lock()
			connCount1++
			mu.Unlock()
			//fmt.Println(atomic.LoadInt32(&connCount))
			//time.Sleep(time.Microsecond * 5)
		}

	}()
	//go func() {
	//	for {
	//		time.Sleep(time.Microsecond * 10)
	//		atomic.AddInt32(&connCount, -1)
	//		fmt.Println(atomic.LoadInt32(&connCount))
	//	}
	//}()

	time.Sleep(time.Second*1)
	fmt.Println(n)
	fmt.Println(connCount1)
}
