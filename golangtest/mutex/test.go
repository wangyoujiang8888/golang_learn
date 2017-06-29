package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//声明
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (G0)")
	//加锁mutex
	mutex.Lock()

	fmt.Println("The lock is locked.(G0)")
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (G%d)\n", i)
		}(i)
	}
	//休息一会,等待打印结果
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (G0)")
	//解锁mutex
	mutex.Unlock()

	fmt.Println("The lock is unlocked. (G0)")
	//休息一会,等待打印结果
	time.Sleep(time.Second)

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
