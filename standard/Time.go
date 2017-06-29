package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
	fmt.Println(time.Now().Sub(t).Seconds())
}
