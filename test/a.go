package main

import (
	"sync/atomic"
	"fmt"
	"time"
)


var connCount int32

func main() {
	atomic.StoreInt32(&connCount, 0)

	fmt.Println(atomic.LoadInt32(&connCount))
	n := 0
	go func() {
		for {
			n++
			//atomic.StoreInt32(&connCount, 0)
			atomic.AddInt32(&connCount, 1)

			//if atomic.LoadInt32(&connCount) > 1000000000 {
			//	return
			//}
			//go func() {
			//	n++
			//	//atomic.StoreInt32(&connCount, 0)
			//	atomic.AddInt32(&connCount, 1)
			//
			//	if atomic.LoadInt32(&connCount) > 1000000000 {
			//		return
			//	}
			//	//atomic.AddInt32(&connCount, -1)
			//	//atomic.LoadInt32(&connCount)
			//	//fmt.Println(atomic.LoadInt32(&connCount))
			//	//time.Sleep(time.Microsecond * 5)
			//}()

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
	fmt.Println(atomic.LoadInt32(&connCount))
}
