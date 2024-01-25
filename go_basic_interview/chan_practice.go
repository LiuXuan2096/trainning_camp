package main

import (
	"fmt"
	"time"
)

func f14() {
	ch := make(chan int, 2)
	ch1 := make(chan int, 2)
	go func() {
		for {
			select {
			case <-ch:
				ch1 <- 1
			}
		}
	}()

	for {
		select {
		case <-ch1:
			ch <- 1
		}
	}
}

func f1() {
	ch := make(chan int)
	close(ch)
	go func() {
		fmt.Println("开始接收数据：")
		fmt.Println("<-ch:", <-ch)
		fmt.Println("数据接收完成")
	}()

	go func() {
		fmt.Println("开始发送数据")
		ch <- 1
		fmt.Println("发送数据完毕")
	}()
	time.Sleep(time.Second)
}

func f2() {
	ch := make(chan int)
	close(ch)
	for {
		select {
		case c := <-ch:
			fmt.Println("读取<-ch: ", c)
			time.Sleep(time.Second)
		}
	}
}

func f3() {
	ch := make(chan int)
	close(ch)
	<-ch
	println("完成") // 如果这句话会输出则说明不会阻塞
}

//func main() {
//	f3()
//}
