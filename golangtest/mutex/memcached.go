package main

import (
	"fmt"
	"sync"
)

func main() {

	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	const TOTAL_WORKERS = 5
	num_workers := 0

	//主线程，在一个goroutine里

	go func() {
		cond.L.Lock()
		for num_workers < TOTAL_WORKERS {
			cond.Wait()
		}
		fmt.Println("all workers up!")
		cond.L.Unlock()

	}()


	go func() {
		//工作线程，同样是在goroutine中
		cond.L.Lock()
		if num_workers++; num_workers >= TOTAL_WORKERS {
			cond.Signal()
		}
		cond.L.Unlock()

	}()

}
