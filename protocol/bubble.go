package main

import "fmt"

func main() {
	list := []int{1, 7, 10, 34, 2, 3, 100}
	BubbleSort(list)

	fmt.Println(list)

}

func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j+1], values[j] = values[j], values[j+1]
				flag = false
			}
		}
		if flag == true {
			break
		}

	}
}



