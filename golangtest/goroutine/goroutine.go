package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, val := range a {
		total += val
	}
	c <- total

}

func fibonacci(n int, c chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		c <- x
		y = x + y
		x = y
	}
	close(c)

}

func cibonacci(add, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case add <- x: //add 中 可以写入的时候触发
			y = x + y
			x = y
		case <-quit: //quit 可以读的时候触发
			fmt.Println("quit")
			return
		default:
			//fmt.Println("没有chanel 满足条件")
		}

	}
}

func write(add, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-add)
	}
	quit <- 0

}

func timeout() {

	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o

}

func Test() {

	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("main thread")

}

func init() {

	//	runtime.GOMAXPROCS(2)
	// go say("world") //开一个新的Goroutines执行
	// say("hello")    //当前Goroutines执行

	// list := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// c := make(chan int)

	// go sum(list[:len(list)/2], c)

	// go sum(list[len(list)/2:], c)

	// x := <-c
	// y := <-c

	// fmt.Println("x+y=", x+y)

	// e := make(chan int, 10)

	// go fibonacci(cap(e), e)

	// for i := range e {
	// 	fmt.Println(i)
	// }

	add := make(chan int)

	quit := make(chan int)

	go write(add, quit)

	cibonacci(add, quit)

	// timeout()

	//Test()

}
