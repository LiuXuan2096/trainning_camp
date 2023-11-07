package main

//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// fun1()
//	// func2()
//	// func3()
//	// func4()
//	// func5()
//	func7()
//}
//
//func fun1() {
//	timer := time.NewTimer(2 * time.Second) // 设置超时时间2s
//	<-timer.C
//	fmt.Println("after 2s Time out!")
//}
//
//func func2() {
//	timer := time.NewTimer(2 * time.Second)
//	res := timer.Stop()
//	fmt.Println(res)
//}
//
//func func3() {
//	timer := time.NewTimer(time.Second * 2)
//	<-timer.C
//	fmt.Println("time out")
//
//	res1 := timer.Stop()
//	fmt.Printf("res1 is %t\n", res1)
//
//	timer.Reset(time.Second * 3)
//
//	res2 := timer.Stop()
//	fmt.Printf("res2 is %t\n", res2)
//}
//
//func func4() {
//	duration := time.Duration(1) * time.Second
//
//	f := func() {
//		fmt.Println("f has been called after 1s by time.AfterFunc")
//	}
//
//	timer := time.AfterFunc(duration, f)
//
//	defer timer.Stop()
//
//	time.Sleep(2 * time.Second)
//}
//
//func func5() {
//	ch := make(chan string)
//
//	go func() {
//		time.Sleep(time.Second * 3)
//		ch <- "葛诗颖"
//	}()
//
//	select {
//	case val := <-ch:
//		fmt.Printf("val is %s\n", val)
//	case <-time.After(time.Second * 2):
//		fmt.Println("timeout!!!")
//
//	}
//}
//
//func func6() chan struct{} {
//	ticker := time.NewTicker(1 * time.Second)
//
//	ch := make(chan struct{})
//	go func(ticker *time.Ticker) {
//		defer ticker.Stop()
//		for {
//			select {
//			case <-ticker.C:
//				fmt.Println("watch!!!")
//			case <-ch:
//				fmt.Println("Ticker stop!!!")
//				return
//
//			}
//		}
//	}(ticker)
//
//	return ch
//}
//
//func func7() {
//	ch := func6()
//	time.Sleep(5 * time.Second)
//	ch <- struct{}{}
//	close(ch)
//}
