package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))

	fmt.Println(strings.ContainsAny("failure", "M & a"))
	fmt.Println(strings.ContainsAny("in failure", "M g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))

	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))

	n, err := strconv.ParseUint("256", 10, 8)

	fmt.Println(n)
	if err != nil {
		fmt.Println(err)
	}

	_, err1 := strconv.Atoi("512")

	if err1 != nil {
		fmt.Println(err1)
	}

	s := "hello wold"
	i := 10

	fmt.Println(s + strconv.FormatInt(int64(i), 10))

	digit := rune('b')
	fmt.Println(digit)

	fmt.Println(unicode.IsDigit(digit))  //true
	fmt.Println(unicode.IsNumber(digit)) //true
	fmt.Println(unicode.IsLower(digit))

}
