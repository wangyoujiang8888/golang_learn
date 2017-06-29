package iota

import (
	"fmt"
)

const (
	x = iota
	y = iota
	z = iota
	w
	k
)

const v = iota
const (
	e, f, g = iota, iota, iota //e=0,f=0,g=0 iota在同一行值相同
)
const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, r, t = iota, iota, iota //d=3,e=3,f=3

)

func Iota() {
	fmt.Println(k)
	fmt.Println(v)
	fmt.Println(g)
	fmt.Println(c)
}
