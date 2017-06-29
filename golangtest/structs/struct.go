package structs

import (
	"fmt"
	"math"
)

type persion struct {
	name string
	age  int
}

type Human struct {
	name   string
	age    int
	weight float32
}

type Student struct {
	Human
	speciality string
}

type Employee struct {
	Human   //匿名字段
	company string
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Color byte

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, deep float64
	color               Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.height * b.width * b.deep
}

func (b *Box) SetColor(color Color) {
	b.color = color
}

func (boxlist BoxList) BiggestColor() Color {

	v := 0.00
	k := Color(WHITE)

	for _, val := range boxlist {
		if volume := val.Volume(); volume > v {
			v = volume
			k = val.color
		}
	}

	return k
}

func (boxlist BoxList) PaintItBlack() {
	for k, _ := range boxlist {
		boxlist[k].SetColor(BLACK)
	}
}

func (color Color) String() string {
	var color_list = [...]string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return color_list[color]
}

func testMethod1() {

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	bigcolor := boxes.BiggestColor()

	fmt.Println("BiggestColor", bigcolor)

	fmt.Println("the biggest color string", bigcolor.String())

	boxes.PaintItBlack()

	for _, val := range boxes {
		fmt.Println("print to new coler string ", val.color.String())
	}

}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func testMethod() {
	r1 := Rectangle{1.2, 2}
	r2 := Rectangle{1.2, 4}

	c1 := Circle{1.2}
	c2 := Circle{2.4}

	fmt.Println("r1 area()", r1.area())

	fmt.Println("r2 area()", r2.area())

	fmt.Println("c1 area()", c1.area())

	fmt.Println("c2 area()", c2.area())

}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s i am age is  %d\n", h.name, h.age)
}

func (h Employee) SayHi() {
	fmt.Printf("Hi, I am %s i am company is  %s\n", h.name, h.company)
}

func (h Student) SayHi() {
	fmt.Printf("Hi, I am %s i study is  %s\n", h.name, h.speciality)
}

func testMethod2() {

	mark := Student{Human{"riverwang", 23, 34.5}, " computer"}
	tom := Employee{Human{"riverwang", 23, 34.5}, "shen chuangxiang chuang xiang tian kong"}

	mark.SayHi()
	tom.SayHi()

	mark.Human.SayHi()
	tom.Human.SayHi()

}

func init() {
	//structs()
	//student()

	//testMethod()
	testMethod1()

	//testMethod2()
}

func student() {
	mark := Student{Human{"Mark", 32, 58.2}, "computers "}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	mark.age += 1

	mark.weight += 1

	fmt.Println("his new age", mark.age)
	fmt.Println("hies new weight", mark.weight)

}
func structs() {

	//直接赋值

	var P = persion{"Tom", 23}
	fmt.Println("p = ", P)

	var P1 = persion{name: "Tom1", age: 24}

	fmt.Println("P1 = ", P1)

	var P2 = new(persion)

	P2.name = "river.wang"

	P2.age = 30

	fmt.Println("P2 = ", P2)

	P12_old, P12_diff := older(P1, *P2)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		P1.name, P2.name, P12_old.name, P12_diff)

}

func older(p1, p2 persion) (persion, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age

}
