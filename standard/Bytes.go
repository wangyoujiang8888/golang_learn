package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 12)

	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, %d\n", p, n)

	//偏移写入
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余的")
	g, err := file.WriteAt([]byte("ggggggggggggggggggggggggggg"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(g)

	ReadFromTest()

	WriterToTest()

	SeekerTest()

	ByteReaderTest()
	fmt.Println("utf8编码问题")

	s := Utf8Index("Go语言中文网", "中文")

	fmt.Println(s)

	LimitedReaderTest()
}

func ReadFromTest() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

func WriterToTest() {
	//reader := bytes.NewReader([]byte("Go语言中文网"))
	//reader := strings.NewReader("Go语言中文网")

	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	reader.WriteTo(os.Stdout)

}

func SeekerTest() {
	reader := strings.NewReader("Go语言中文网")

	reader.Seek(-6, os.SEEK_END)

	/*r, _ := reader.ReadByte()

	fmt.Printf("%d\n", r)

	g, _ := reader.ReadByte()

	fmt.Printf("%d\n", g)

	h, _ := reader.ReadByte()

	fmt.Printf("%d\n", h)
	*/
	m, _, _ := reader.ReadRune()

	fmt.Println(m)

}

func ByteReaderTest() {
	var ch byte
	fmt.Scanf("%c\n", &ch)
	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)

	if err == nil {
		fmt.Println("写入一个字节成功！准备读取该字节……")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读取的字节：%c\n", newCh)
	} else {
		fmt.Println("写入错误")
	}
}

// 即 Utf8Index("Go语言中文网", "中文") 返回 4，而不是 strings.Index 的 8
func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for rune_c, size, err := reader.ReadRune(); err == nil; rune_c, size, err = reader.ReadRune() {
		fmt.Println(rune_c)
		fmt.Println(size)
		totalSize += size
		pos++
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos

}

func LimitedReaderTest() {
	content := "This Is LimitReader Example"
	reader := strings.NewReader(content)
	limitReader := &io.LimitedReader{R: reader, N: 8}
	for limitReader.N > 0 {
		tmp := make([]byte, 2)
		limitReader.Read(tmp)
		fmt.Printf("%s", tmp)
	}

}
