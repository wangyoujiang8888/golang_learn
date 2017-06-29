package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{"Newmarch", "Jan"},
		Email: []Email{Email{"home", "346006742@qq.com"},
			Email{"work", "wangyoujiang8888@163.com"}}}

	saveGob("person.gob", person)

}

func saveGob(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	decoder := gob.NewEncoder(outFile)
	err = decoder.Encode(key)
	checkError(err)
	defer outFile.Close()

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

}
