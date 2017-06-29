package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	Rsn() int64
	// 获取最后写入的数据块的序列号
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}
type Data []byte

//数据文件的实现类型
type myDataFile struct {
	f       *os.File     //文件
	fmutex  sync.RWMutex //被用于文件的读写锁
	woffset int64        // 写操作需要用到的偏移量
	roffset int64        // 读操作需要用到的偏移量
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁
	dataLen uint32       //数据块长度
	rcond   *sync.Cond
}

//初始化DataFile类型值的函数,返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	os.Remove(path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	//f,err := os.Create(path)
	if err != nil {
		fmt.Println("Fail to find", f, "cServer start Failed")
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f:       f,
		dataLen: dataLen,
	}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.rcond.Wait()
				continue
			}
		}
		break
	}
	d = bytes
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()
	wsn = offset / int64(df.dataLen)
	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	df.rcond.Broadcast()
	return
}

func (df *myDataFile) Rsn() (rsn int64) {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	rsn = df.roffset
	return
}

func (df *myDataFile) Wsn() (wsn int64) {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	wsn = df.woffset
	return
}

func (df *myDataFile) DataLen() (d uint32) {
	d = uint32(df.dataLen)
	return
}

func main() {
	var dataFile DataFile
	dataFile, _ = NewDataFile("./mutex_2015_2.dat", 10)
	var d = map[int]Data{
		1: []byte("batu_test1"),
		2: []byte("batu_test2"),
		3: []byte("batu_test3"),
		4: []byte("batu_test4"),
		5: []byte("batu_test5"),
		6: []byte("batu_test6"),
		7: []byte("batu_test7"),
		8: []byte("batu_test8"),
		9: []byte("batu_test9"),
	}
	for i := 1; i < 10; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			fmt.Println("write i=", i, ",wsn=", wsn, ",success.")
		}(i)
	}
	for i := 1; i < 10; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read()
			fmt.Println("Read i=", i, ",rsn=", rsn, ",data=", d, ",success.")
		}(i)
	}
	time.Sleep(10 * time.Second)
}
