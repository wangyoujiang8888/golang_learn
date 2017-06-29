package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	bb := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, bb)
	encoder.Write(eightBitData)
	encoder.Close()

	fmt.Println(bb)

	decoder := base64.NewDecoder(base64.StdEncoding, bb)
	dbuf := make([]byte, 12)

	decoder.Read(dbuf)

	for _, v := range dbuf {
		fmt.Println(v)
	}

	gg := &bytes.Buffer{}
	encoder = base64.NewEncoder(base64.StdEncoding, gg)

	tt := "dddddddddddddddd"
	encoder.Write([]byte(tt))

	fmt.Println(gg)

	hh := make([]byte, 20)

	decoder = base64.NewDecoder(base64.StdEncoding, gg)

	decoder.Read(hh)

	fmt.Println("content:" + string(hh))

}
