package slice

import (
	"fmt"
)

func Slice1() {
	var ar = []byte{'a', 'b', 'c', 'e', 'f', 'g'}
	var a, b []byte
	a = ar[3:]
	b = ar[:]

	fmt.Printf("%c", a[0])
	fmt.Printf("%c", a[2])

	fmt.Printf("%c", b[0])
	fmt.Printf("%c", b[1])

	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	var Slice_b []byte

	copy(Slice_b, Slice_a)

	fmt.Printf("\nlen=%d", len(Slice_a))
	fmt.Printf("\ncap=%d", cap(Slice_a))

	fmt.Printf("Slice_b len=%d", len(Slice_b))

}
