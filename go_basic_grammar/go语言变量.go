package main

/*
全局变量的生命周期是程序存活时间
在不发生内存逃逸的情况下，局部变量的生命周期是函数存活时间
*/

//func main() {
//	// func1()
//	// func2()
//	// func3()
//	// func4()
//	// func5()
//	// func6()
//	func7()
//}

// 这种因式分解关键字的写法一般用于声明全局变量
// 全局变量允许声明但不适用
//var (
//	a int
//	b bool
//)
//
//func func1() {
//	var a string = "葛诗颖"
//	fmt.Println(a)
//
//	var b int
//	fmt.Println(b)
//
//	var c bool
//	fmt.Println("bool变量c的零值为：", c)
//}
//
//func func2() {
//	// 以下变量的零值均为nil
//	var a *int
//	var b []int
//	var c map[string]int
//	var d chan int
//	var e func(string) int
//	var f error
//
//	fmt.Println("Hello World!")
//	fmt.Println(a == nil)
//	fmt.Println(b == nil)
//	fmt.Println(c == nil)
//	fmt.Println(d == nil)
//	fmt.Println(e == nil)
//	fmt.Println(f == nil)
//}
//
//func func3() {
//	// 测试以下类型变量的零值
//	var i int
//	var f float64
//	var b bool
//	var s string
//	fmt.Printf("%v %v %v %q\n", i, f, b, s)
//}
//
//func func4() {
//	// 根据值自动判断变量类型
//	var d = true
//	fmt.Printf("变量d的类型是%T", d)
//}
//
//func func5() {
//	// :=声明变量
//	// 如果变量已经使用var声明过了，再使用 := 声明变量，会产生编译错误
//	girl := "葛诗颖"
//	fmt.Println(girl)
//}
//
//func func6() {
//	var x, y int
//	var c, d int = 1, 2
//	var e, f = 123, "葛诗颖"
//	// 这种不带声明格式的只能在函数体中出现
//	g, h := 123, true
//	fmt.Println(a, b, x, y, c, d, e, f, g, h)
//}
//
//func func7() {
//	// 交换两个变量a，b， 两个变量的类型必须相同
//	girl1, girl2 := "葛诗颖", "GeShiYing"
//	girl1, girl2 = girl2, girl1
//	fmt.Println(girl1, girl2)
//	// _ 空白标识符，通常被用于抛弃值，是一个只写变量，你不能得到它的值
//	_, c := 5, 7
//	fmt.Println(c)
//}
