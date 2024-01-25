package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func main() {
	// func1()
	// func2()
	// func3()
	// func5()
	// func4()
	// func6()
	func7()
}

func func1() {
	ss := "葛诗颖"
	for _, v := range ss {
		fmt.Printf("%c: %d\n", v, v)
	}
}

func func2() {
	var ss string
	ss = "葛诗颖"
	strByte := []rune(ss)
	strByte[2] = 35799
	fmt.Println(string(strByte))
}

func func3() {
	a := []byte{1, 2, 3}
	// 构建string.Builder 和 bytes.Buffer
	b := strings.Builder{}
	b.Write(a)
	b2 := bytes.NewBuffer(a)
	str1 := b.String()
	str2 := b.String()
	// 通过反射来获取str1 str2的底层描述
	String2Bytes(str1)
	String2Bytes(str2)

	str3 := b2.String()
	str4 := b2.String()
	String2Bytes(str3)
	String2Bytes(str4)
}

func String2Bytes(s string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	fmt.Println(bh.Data)
}

func func5() {
	arr1 := make([]int, 3, 5)
	arr2 := make([]int, 3)
	fmt.Printf("arr1,len=%d, cap=%d\n", len(arr1), cap(arr1))
	fmt.Printf("arr2,len=%d, cap=%d\n", len(arr2), cap(arr2))

	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := arr3[1:3]
	fmt.Printf("arr3=%v\n", arr3)
	fmt.Printf("arr4=%v\n", arr4)
}

func func4() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr2[0] = 100
	fmt.Printf("arr1=%v\n", arr1)
	fmt.Printf("arr2=%v\n", arr2)

	arr3 := [3]int{1, 2, 3}
	arr4 := arr3
	arr4[0] = 100
	fmt.Printf("arr3=%v\n", arr3)
	fmt.Printf("arr4=%v\n", arr4)
}

func func6() {
	arr := []int{1, 2, 3, 4}
	arr1 := make([]int, 3)
	cnt := copy(arr1, arr)
	fmt.Printf("cnt=%d\n", cnt)
	fmt.Printf("arr1=%v\n", arr1)
}

type people struct {
	name string
}

var girl = people{
	name: "葛诗颖",
}

func printPeople(u <-chan *people) {
	time.Sleep(2 * time.Second)
	fmt.Println("printPeople", <-u)
}

func func7() {
	c := make(chan *people, 5)
	var a = &girl

	c <- a
	fmt.Println(a)
	a = &people{
		name: "诗诗",
	}

	go printPeople(c)
	time.Sleep(5 * time.Second)
	fmt.Println(a)
}
