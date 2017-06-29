package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name    string
	Email   string
	Tel     string
	Address string
}

type Msg struct {
	Ret    int
	msg    string
	Person Person
}

func (p Person) String() string {
	s := p.Name + " " + p.Email + "" + p.Tel + "" + p.Address
	return s
}

func main() {
	person1 := Person{Name: "riverwang", Email: "346006742@qq.com", Tel: "18566221038", Address: "ttt"}

	list1 := list.New()
	list1.PushBack(person1)

	person2 := Person{Name: "riverwang1", Email: "346006742@qq.com", Tel: "18566221038", Address: "ttt"}

	list1.PushBack(person2)

	service := ":1201"

	listener, err := net.Listen("tcp", service)

	checkError(err)
	var person Person
	var msg Msg

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		decoder := json.NewDecoder(conn)
		decoder.Decode(&person)

		encoder := json.NewEncoder(conn)

		for e := list1.Front(); e != nil; e = e.Next() {
			if value, ok := e.Value.(Person); ok {
				if value.Name == person.Name && value.Email == person.Email {
					fmt.Println("找到用户" + value.String())
					msg = Msg{0, "成功", value}
					encoder.Encode(msg)

					break
				} else {
					msg = Msg{-1, "失败", person1}
					encoder.Encode(msg)
					fmt.Println("未找到用户信息")
					break
				}
			}
		}

		//fmt.Println(person.String())

		conn.Close()

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
