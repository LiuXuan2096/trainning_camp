package main

//import (
//	"fmt"
//	"time"
//)
//
//type SChannel = chan<- int // 定义写channel
//type RChannel = <-chan int // 定义读channel
//
//func main() {
//	// func1()
//	// func2()
//	// func3()
//	// func4()
//	// func5()
//	ch := make(chan bool, 1)
//
//	var num int
//	for i := 0; i < 100; i++ {
//		go add(ch, &num)
//	}
//	time.Sleep(3 * time.Second)
//	fmt.Println("num 的值:", num)
//}
//
//func func1() {
//	ch := make(chan int, 5)
//	ch <- 1
//	close(ch)
//	go func() {
//		for i := 0; i < 5; i++ {
//			v := <-ch
//			fmt.Printf("v=%d\n", v)
//		}
//	}()
//	// time.Sleep(2 * time.Second)
//}
//
//func func2() {
//	ch := make(chan int, 5)
//	ch <- 1
//	close(ch)
//	go func() {
//		for i := 0; i < 5; i++ {
//			v, ok := <-ch // 判断句式读取
//			if ok {
//				fmt.Printf("v = %d\n", v)
//			} else {
//				fmt.Printf("channel数据已读完, v= %d\n", v)
//			}
//		}
//	}()
//}
//
//func func3() {
//	ch := make(chan int, 5)
//	ch <- 1
//	ch <- 2
//	close(ch)
//	go func() {
//		for v := range ch {
//			fmt.Printf("v=%d\n", v)
//		}
//	}()
//
//}
//
//func func4() {
//	var ch = make(chan int)
//
//	go func() {
//		var send SChannel = ch
//		fmt.Println("send: 100")
//		send <- 100
//	}()
//
//	go func() {
//		var rec RChannel = ch
//		num := <-rec
//		fmt.Printf("receive: %d", num)
//	}()
//}
//
//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum
//}
//
//func func5() {
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go func() {
//		sum(s[:len(s)>>1], c)
//	}()
//	go sum(s[len(s)>>1:], c)
//
//	x, y := <-c, <-c
//	fmt.Println(x, y, x+y)
//}
//
//func add(ch chan bool, num *int) {
//	ch <- true
//	*num = *num + 1
//	<-ch
//}
