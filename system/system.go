package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

func storedata(data interface{}, filename string) {
	//buffer为bytes.Buffer结构，这个结构实际上就是一个拥有Read和Write方法的可变长度字节缓冲区，即bytes.Buffer既是读取器也是写入器
	mybuffer := new(bytes.Buffer)

	//将缓冲区数据传递给NewEncoder函数，以此来创建出一个gob编码器
	encode := gob.NewEncoder(mybuffer)

	//调用编码器的Encode函数，将数据data写入到mybuffer中
	err := encoder.Encode(data)

	//将mybuffer数据写入到文件中
	err = ioutil.WriteFile(filename, mybuffer.Bytes(), 0600)
}

func loaddata(data interface{}, filename string) {
	//读取文件数据存入到原始数据raw中
	raw, err := ioutil.ReadFile(filename)

	//根据原始数据raw创建缓冲区buffer，借此为原始数据raw提供相应的Read方法和Write方法
	buffer := bytes.NewBuffer(raw)

	//调用NewDecoder函数，为缓冲区创建相应的解码器
	dec := gob.NewDecoder(buffer)

	//用解码器读取原始数据，并存入到data中
	err = dec.Decode(data)
}

func main() {
	//存储数据post到post1文件中：

	storedata(post, "post1")
	//读取文件post1数据到post中：

	loaddata(&post, "post1")
}
