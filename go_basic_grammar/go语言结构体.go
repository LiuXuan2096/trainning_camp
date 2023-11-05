package main

//import "fmt"
//
//type Student struct {
//	ID    int
//	Name  string
//	Age   int
//	Score int
//}
//
//func main() {
//	// func1()
//	// func2()
//	func3()
//}
//
//func func1() {
//	// 键值对方式初始化结构体
//	st := Student{
//		ID:    100,
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 100,
//	}
//	fmt.Printf("学生st:%v\n", st)
//}
//
//func func2() {
//	// 值列表初始化
//	// 以值列表的方式初始化结构体，值列表的个数必须等于结构体属性的个数，
//	// 且要按顺序，否则会报错
//	st := Student{
//		100,
//		"葛诗颖",
//		22,
//		100,
//	}
//	fmt.Printf("葛诗颖:%v\n", st)
//}
//
///*
//使用 . 操作符来访问结构体的成员，. 前可以是结构体变量或者结构体指针
//*/
//func func3() {
//	girl := Student{
//		ID:    100,
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 100,
//	}
//	fmt.Printf("女生的姓名是:%s\n", girl.Name)
//
//	pgirl := &girl
//	fmt.Printf("女生的年龄是:%d\n", pgirl.Age)
//}
