package main

import "fmt"

type Website struct {
	Name string
}

var site = Website{Name: "studygolang"}

func main() {
	fmt.Printf("%v", site)
	fmt.Printf("%+v", site)

	fmt.Printf("%#v", site)
	fmt.Printf("%T", site)

	fmt.Printf("%t", true)

	fmt.Printf("%b", 5)

}
