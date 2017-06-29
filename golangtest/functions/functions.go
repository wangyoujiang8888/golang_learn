package functions

import (
	f "fmt"
)

type TestInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f TestInt) []int {

	var result []int
	for _, val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

func TestFuctnion() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	odd := filter(a, isOdd)
	f.Println("Odd elements of slice are: ", odd)

	even := filter(a, isEven)
	f.Println("even elements of slice are: ", even)

}

func Eend() {
	f.Println("Eend functions.go")
}

func init() {
	TestFuctnion()
}
