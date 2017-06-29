package arr

import (
	"fmt"
)

func Aarr1() {
	var a [10]int32
	a[0] = 100
	a[1] = 200
	fmt.Println(a[1])

	b := [10]int{1, 2, 3, 4}
	c := [...]int{6, 7, 8, 9}

	fmt.Println(b[0])
	fmt.Println(c[0])

	doubleArray := [2][4]int{[4]int{1, 2}, [4]int{3, 4}}

	fmt.Println(doubleArray[0][1])

	threeArray := [2][4]int{{1, 2}, {3, 4}}

	fmt.Println(threeArray[0][1])
}
