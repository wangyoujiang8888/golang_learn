package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	locker := new(sync.Mutex)
	ch := make(chan int)
	go func() {
		locker.Lock()
		fmt.Println("get lock first")
		time.Sleep(10 * time.Second)
		locker.Unlock()
	}()

	go func() {
		locker.Lock()
		fmt.Println("hello, lock")
		locker.Unlock()
		ch <- 1
	}()

	fmt.Println("main")
	<-ch //主线程等待完成
}
