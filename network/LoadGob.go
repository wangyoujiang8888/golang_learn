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

func (p Person) String() string {
	s := p.Name.Family + " " + p.Name.Personal + "\n"
	for _, v := range p.Email {
		s += v.Kind + " " + v.Address + "\n"
	}
	return s
}

func main() {
	var p Person
	loadGob("person.gob", &p)
	fmt.Println(p)

}

func loadGob(filename string, key interface{}) {
	outfile, err := os.Open(filename)
	checkError(err)
	decoder := gob.NewDecoder(outfile)
	err = decoder.Decode(key)
	checkError(err)
	defer outfile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
