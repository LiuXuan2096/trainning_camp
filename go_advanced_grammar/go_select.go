package main

//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// func1()
//	// func2()
//	// func3()
//	func4()
//}
//
//func func1() {
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//
//	select {
//	case <-ch1:
//		fmt.Printf("received from ch1")
//	case num := <-ch2:
//		fmt.Printf("num is: %d", num)
//
//	}
//}
//
//func func2() {
//	ch1 := make(chan int, 1)
//
//	select {
//	case <-ch1:
//		fmt.Printf("received from ch1")
//	default:
//		fmt.Println("default!!!")
//
//	}
//}
//
//func func3() {
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	go func() {
//		time.Sleep(time.Second)
//		for i := 0; i < 3; i++ {
//			select {
//			case v := <-ch1:
//				fmt.Printf("Received from ch1, val = %d\n", v)
//			case v := <-ch2:
//				fmt.Printf("Received from ch2, val = %d\n", v)
//			default:
//				fmt.Println("default!!!")
//			}
//			time.Sleep(time.Second)
//		}
//	}()
//	ch1 <- 1
//	time.Sleep(time.Second)
//	ch2 <- 2
//	time.Sleep(4 * time.Second)
//}
//
//func func4() {
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	ch1 <- 5
//	ch2 <- 6
//	select {
//	case v := <-ch1:
//		fmt.Printf("Received from ch1, val = %d\n", v)
//	case v := <-ch2:
//		fmt.Printf("Received from ch2, val = %d\n", v)
//	default:
//		fmt.Println("default!!!")
//	}
//}
