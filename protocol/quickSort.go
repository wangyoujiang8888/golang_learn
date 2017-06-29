package main

import (
	"fmt"
)

func main() {
	values := []int{72, 6, 57, 88, 60, 42, 83, 73, 48, 85}
	quickSort2(values, 0, len(values)-1)
	fmt.Println(values)
	fmt.Println("ggg")
	values1 := []int{72, 6, 57, 88, 60, 42, 83, 73, 48, 85}
	quickSort(values1, 0, len(values1)-1)

	//fmt.Println(values1)

}

func quickSort(values []int, left int, right int) {

	if left < right {
		i := AdjustArray(values, left, right)
		//fmt.Println(i)
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}

func AdjustArray(values []int, left int, right int) int {
	i := left
	j := right
	x := values[i]

	for i < j {
		for i < j && values[j] >= x {
			j--
		}

		if i < j {
			values[i] = values[j]
			i++
		}

		for i < j && values[i] < x {
			i++
		}
		if i < j {
			values[j] = values[i]
			j--
		}
	}
	values[i] = x
	fmt.Println(values)
	return i
}

func quickSort2(values []int, left, right int) {
	if left < right {
		i := left
		j := right
		x := values[(left+right)/2]

		for {
			if values[i] < x {
				i++
			}
			if values[j] > x {
				j--
			}
			if i >= j {
				break
			}
			values[i], values[j] = values[j], values[i]
		}

		quickSort2(values, left, i)
		quickSort2(values, i+1, right)

	}

}
