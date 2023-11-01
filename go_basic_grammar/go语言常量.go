package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// func1()
	// func2()
	// func3()
	func4()
}

func func1() {
	// 常量中的数据类型只可以是布尔型，数字型（整数型、浮点型和复数）和字符串型
	const b string = "葛诗颖"
	const c = "葛诗颖"
	fmt.Println(b, c)
}

func func2() {
	// go语言不像C++ java一样有专门的枚举类型，go语言枚举一般用常量表示
	const Unknown, Success, Fail = 0, 1, 2
	const girl = "葛诗颖"
	const c = unsafe.Sizeof(girl)
	fmt.Println(girl, c)
	/*
		"葛诗颖"是三个汉字，三个汉字大小应为12个字节，但为什么c的值为16？
		实际上 字符串类型的unsafe.Sizeof()一直是16.因为字符串类型对应一个结构体，该结构体有两个域。
		第一个域是指向该字符串的指针，第二个域是该字符串的长度，每个域占8个字节。但是并不包含指针指向
		字符串的内容，这就是为什么sizeof始终返回的是16
	*/
}

/*
iota是一个特殊的常量，可以被认为是一个可以被编译器修改的常量。iota在const关键字出现时会被重置为0，在const内部
的第一行之前，const中每新增一行常量声明将使iota计数一次（iota可理解为const语句块中的行索引）
*/
func func3() {
	const (
		a = iota // 0
		b
		c
		d = "葛诗颖" // 独立值，iota += 1
		e         // iota += 1
		f = 100   // iota += 1
		g         // iota += 1
		h = iota  // iota恢复计数
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

/*
关于iota的一个有趣的例子
*/
func func4() {
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Printf("i=%v, j=%v, k=%v, l=%v", i, j, k, l)
}
