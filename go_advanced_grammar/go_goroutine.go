package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func myGroutine() {
//	fmt.Println("myGroutine")
//}
//
//func main() {
//	//go myGroutine()
//	//fmt.Println("end!!!")
//	//time.Sleep(2 * time.Second)
//
//	//var wg sync.WaitGroup
//	//wg.Add(2)
//	//
//	//go myGoroutine("goroutine1", &wg)
//	//go myGoroutine("goroutine2", &wg)
//	//
//	//wg.Wait()
//
//	//go func() {
//	//	defer func() {
//	//		if e := recover(); e != nil {
//	//			fmt.Printf("sub1 recover:%v\n", e)
//	//		}
//	//	}()
//	//	panic("sub1 func panic!!!")
//	//	fmt.Println("111")
//	//}()
//	//
//	//go func() {
//	//	defer func() {
//	//		if e := recover(); e != nil {
//	//			fmt.Printf("sub2 recover:%v\n", e)
//	//		}
//	//	}()
//	//	panic("sub2 func panic!!!")
//	//	fmt.Println("111")
//	//}()
//	//
//	//fmt.Println("222") // 发生panic后 不会打印
//	//time.Sleep(2 * time.Second)
//
//	panicTest1()
//	panic("main panic")
//	fmt.Println(222222)
//	time.Sleep(2 * time.Second)
//
//}
//
//func myGoroutine(name string, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 0; i < 5; i++ {
//		fmt.Printf("myGoroutine %s\n", name)
//		time.Sleep(10 * time.Millisecond)
//	}
//}
//
//// 一个recover只能捕获一次panic，且一一对应
//func panicTest() {
//	panic("panicTest panic")
//	fmt.Println(111111)
//}
//
//func panicTest1() {
//	defer func() {
//		if e := recover(); e != nil {
//			fmt.Printf("panicTest1 recover:%v\n", e)
//		}
//	}()
//	panicTest()
//}
