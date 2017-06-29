package interfaces

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (h Human) String() string {
	return "❰" + h.name + h.phone + "❱"
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func init() {

	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	// tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = paul

	fmt.Println("This is Paul , a Student")

	i.SayHi()

	i.Sing("精忠报国")

	i = sam

	fmt.Println("This is Sam, a Employee:")

	i.SayHi()

	i.Sing("如果你要离去")

	// //定义数组实现
	// var x [4]Men

	// x[0] = mike
	// x[1] = paul
	// x[2] = sam
	// x[3] = tom

	// for _, val := range x {
	// 	//val.SayHi()
	// }

	//定义 slice

	y := make([]Men, 3)

	y[0], y[1], y[2] = paul, sam, mike

	for _, val := range y {
		val.SayHi()
	}

	Bob := Human{"Bob", 39, "000-7777-XXX"}

	fmt.Println("This Human is : ", Bob)

	assert1()

}

type Element interface{}

type Box []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "my am is " + p.name + "name is " + p.name
}
func assert1() {

	list := make(Box, 3)

	list[0] = 100
	list[1] = "hello worlds"
	list[2] = Person{"river.wang", 30}

	for index, element := range list {

		if val, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, val)
		} else if val, ok := element.(string); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, val)
		} else if val, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, val)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}

	}

	for index, element := range list {

		switch val := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, val)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, val)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, val)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}

	}

}
