/*

一，什么是深浅拷贝：
1）深拷贝（Deep Copy）：
拷贝的数据，拷贝时创建一个新对象，开辟一个新的内存空间，把原对象的数据复制过来，新对象修改数据时不会影响原对象的值。既然数据内存地址不同，释放内存地址时，需要分别释放。
* 值类型的数据，默认都是深拷贝 int ，float，string，bool，array，struct


2）浅拷贝（Shallow Copy）：
拷贝的是数据地址，拷贝时创建一个新对象，然后复制指向的对象的指针。此时新对象和原对象指向的地址都是一样的，因此，新对象修改数组时，会影响原来对象。
引用类型的数据，默认都是浅拷贝 slice，map

二，深浅拷贝转换
 可以使用Encode函数

*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type staff struct {
	Name string
	Age  int32
}

func Clone(src, dst interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(src); err != nil {
		return err
	}
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

func deepCopy() {
	fmt.Println("deep copy 内容一样，地址不同，修改其中一值，内容不变化，另一个不变化 ")
	staff1 := staff{
		Name: "立峰",
		Age:  3,
	}
	staff2 := staff1
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, &staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, &staff2)

	fmt.Println("修改stall1的age属性")
	staff1.Age = 1
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, &staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, &staff2)
}

func deepCopyTurn() {
	fmt.Println("deep copy 转换为 shall copy ")
	staff1 := staff{
		Name: "立峰",
		Age:  3,
	}
	staff2 := &staff1
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, &staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, &staff2)

	fmt.Println("修改stall1的age属性")
	staff1.Age = 1
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, &staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, &staff2)
}

func shallowCopy() {
	fmt.Println("shall copy 内容一样，地址一样，修改其中一值，内容变化，另一个变化 ")
	staff1 := &staff{}
	staff1.Name = "qlf"
	staff1.Age = 3

	staff2 := staff1
	fmt.Println(staff2 == staff1)
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, staff2)

	fmt.Println("修改stall1的age属性")
	staff1.Age = 1
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, staff2)
}

func shallowCopyTurn() {
	fmt.Println("shall copy 转换为 Deep copy ")
	staff1 := new(staff)
	staff1.Name = "qlf"
	staff1.Age = 3

	staff2 := &staff{}
	_ = Clone(staff1, staff2)
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, staff2)
	fmt.Println("修改stall1的age属性")

	staff1.Age = 1
	fmt.Printf("stall1: %+v\t内存地址:%p\n", staff1, staff1)
	fmt.Printf("stall2: %+v\t内存地址:%p\n", staff2, staff2)
}

func main() {
	deepCopy()
	fmt.Println("")
	deepCopyTurn()
	fmt.Println("")
	shallowCopy()
	fmt.Println("")
	shallowCopyTurn()

}
