package main

//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func defer1() {
//	fmt.Println("defer1")
//}
//
//func defer2() {
//	fmt.Println("defer2")
//}
//
//func defer3() {
//	fmt.Println("defer3")
//}
//
//func main() {
//	// func1()
//	//_, err := copyFile(".\\go_basic_grammar\\葛诗颖.txt", "D:\\GoProject\\训练营\\go_basic_grammar\\test.txt")
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//}
//	// func2()
//	// func3()
//	//func4()
//	res := func5()
//	fmt.Println(res)
//}
//
//func func1() {
//	defer defer1()
//	defer defer2()
//	defer defer3()
//}
//
///*
//defer关键字一般用在以下两个场景中：
//1.资源的释放
//2.配合recover一起处理panic
//*/
//func copyFile(dstFile, srcFile string) (wr int64, err error) {
//	src, err := os.Open(srcFile)
//	if err != nil {
//		return
//	}
//	defer src.Close()
//
//	dst, err := os.Create(dstFile)
//	if err != nil {
//		return
//	}
//	defer dst.Close()
//
//	wr, err = io.Copy(dst, src)
//	return wr, err
//}
//
//func func2() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println(r)
//		}
//	}()
//	a := 1
//	b := 0
//	fmt.Println("result:", a/b)
//}
//
//func func3() {
//	var num = 1
//	defer fmt.Printf("num is %d", num)
//
//	num = 2
//	return
//}
//
//func func4() {
//	var arr = [4]int{1, 2, 3, 4}
//	defer printArr(&arr)
//
//	arr[0] = 100
//	return
//}
//
//func printArr(arr *[4]int) {
//	for i := range arr {
//		fmt.Println(arr[i])
//	}
//}
//
//func func5() (res int) {
//	num := 1
//
//	defer func() {
//		res++
//	}()
//
//	return num
//}
//
//func func6() int {
//	var num int
//	defer func() {
//		num++
//	}()
//
//	return 1
//}
