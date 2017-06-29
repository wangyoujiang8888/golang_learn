package mmap

import (
	"fmt"
)

func Mmap() {

	numbers := make(map[string]int)

	numbers["one"] = 1

	numbers["two"] = 2

	fmt.Println("numbers key one = ", numbers["one"])

	rating := map[string]float32{"c": 2.1, "c#": 3.0, "php": 5.0, "golang": 6.0}

	csharpRating, ok := rating["c#"]

	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了

	fmt.Println("m key Hello = ", m["Hello"])

	for i := 0; i < 10; i++ {

		for index := 10; index > 0; index-- {
			if index == 5 {
				break
			}
			fmt.Println("index = ", index)
		}

		fmt.Println("i = ", i)
	}

	for k, v := range m {
		fmt.Println("key = ", k)
		fmt.Println("val =", v)
	}

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}
