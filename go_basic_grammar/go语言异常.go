package main

//import "fmt"
//
//func main() {
//	func2()
//}
//
//func fun1() {
//	defer func() {
//		if error := recover(); error != nil {
//			fmt.Println("出现了panic 使用recover获取信息:", error)
//		}
//	}()
//
//	fmt.Println("1111111111")
//	panic("出现panic")
//	fmt.Println("22222222")
//}
//
//func testPanic() {
//	fmt.Println("testPanic上半部分")
//	testPanic2()
//	fmt.Println("testPanic下半部分")
//}
//
//func testPanic2() {
//	defer func() {
//		recover()
//	}()
//	fmt.Println("testPanic2上半部分")
//	testPainc3()
//	fmt.Println("testPanic2下半部分")
//}
//
//func testPainc3() {
//	fmt.Println("testPanic3上半部分")
//	panic("在testPanic3出现了panic")
//	fmt.Println("testPanic3下半部分")
//}
//
//func func2() {
//	fmt.Println("程序开始")
//	testPanic()
//	fmt.Println("程序结束")
//}
