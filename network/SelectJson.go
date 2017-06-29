package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type Person struct {
	Name  string
	Email string
}

type Msg struct {
	Ret    int
	msg    string
	Person Person
}

func (p Msg) String() string {
	s := p.msg + " " + p.Person.Name + " " + p.Person.Email
	return s
}

func main() {
	arg_num := len(os.Args)
	if arg_num < 3 {
		fmt.Print(
			"Please runAs [super user] in [terminal].\n",
			"Usage:\n",
			"\tSelectJson name email\n",
			"\texample: SelectJson riverwang 346006742@qq.com",
		)
		time.Sleep(5e9)
		return
	}

	person := Person{os.Args[1], os.Args[2]}
	conn, err := net.Dial("tcp", ":1201")
	checkError(err)
	encoder := json.NewEncoder(conn)
	encoder.Encode(person)

	decoder := json.NewDecoder(conn)
	var msg Msg
	decoder.Decode(&msg)
	fmt.Println("返回数据" + msg.String())
	defer conn.Close()

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
