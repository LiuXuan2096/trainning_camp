package main

//import (
//	"fmt"
//	"math"
//)
//
//// 声明一个函数类型
//type fc func(int) int
//
//func main() {
//	// func2()
//	// func3()
//	// func4()
//	// func5()
//	// func6()
//	func7()
//}
//
//func func1(x, y string) (string, string) {
//	return y, x
//}
//
//func func2() {
//	boy, girl := func1("葛诗颖", "刘轩")
//	fmt.Println(boy, " falls in love with ", girl)
//}
//
//func func3() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200
//	var ret int
//
//	// 调用函数返回最大值
//	ret = max(a, b)
//	fmt.Printf("最大值:%d\n", ret)
//}
//
//func max(num1, num2 int) int {
//	if num1 > num2 {
//		return num1
//	} else {
//		return num2
//	}
//}
//
///*
//在golang中函数的参数传递都是值传递，不存在引用传递（区别于C++）
//注意：使用值传递时，虽然会改变形参的值，但不会改变函数外变量即实参的值
//*/
//func func4() {
//	var a int = 100
//
//	fmt.Printf("自增前 a的值为:%d\n", a)
//	add(a)
//	fmt.Printf("自增后 a的值为:%d\n", a)
//}
//
//func add(a int) {
//	a++
//	fmt.Printf("add里a的值:%d\n", a)
//}
//
//func func5() {
//	// 声明函数变量
//	getSquareRoot := func(x float64) float64 {
//		return math.Sqrt(x) // 求一个数的平方根
//	}
//
//	// 使用函数变量调用函数
//	fmt.Println(getSquareRoot(9))
//}
//
//func func6() {
//	CallBack(1, callBack)
//}
//
//func CallBack(x int, f fc) {
//	f(x)
//}
//
//func callBack(x int) int {
//	fmt.Printf("我是回调, x:%d\n", x)
//	return x
//}
//
//func func7() {
//	// nextNumber 为一个函数 函数中i为0
//	nextNumber := getNumber()
//
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//
//	// 创建新的函数nextNumber1, 并查看结果
//	nextNumber1 := getNumber()
//	fmt.Println(nextNumber1())
//	fmt.Println(nextNumber1())
//}
//
//func getNumber() func() int {
//	i := 0
//	return func() int {
//		i += 1
//		return i
//	}
//}
