package main

//import (
//	"context"
//	"fmt"
//	"time"
//)
//
//func main() {
//	// func1()
//	// func2()
//	// func3()
//	func5()
//}
//
//func func1() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go Watch(ctx, "goroutine1")
//	go Watch(ctx, "goroutine2")
//
//	time.Sleep(6 * time.Second)
//	fmt.Println("end watching!!!")
//	cancel()
//	time.Sleep(1 * time.Second)
//}
//
//func Watch(ctx context.Context, name string) {
//	for {
//		select {
//		case <-ctx.Done():
//			// 主goroutine调用cancel后，会发送一个信号到ctx.Done()这个channel，这里就会收到信息
//			fmt.Printf("%s exit!\n", name)
//			return
//		default:
//			fmt.Printf("%s watching...\n", name)
//			time.Sleep(time.Second)
//		}
//	}
//}
//
//func func2() {
//	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
//	defer cancel()
//	go Watch(ctx, "goroutine1")
//	go Watch(ctx, "goroutine2")
//
//	time.Sleep(6 * time.Second)
//	fmt.Println("end watching!!!")
//}
//
//func func3() {
//	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
//	defer cancel()
//	go Watch(ctx, "goroutine1")
//	go Watch(ctx, "goroutine2")
//
//	time.Sleep(6 * time.Second)
//	fmt.Println("end watching!!!")
//}
//
//func func4(ctx context.Context) {
//	fmt.Printf("name is: %s", ctx.Value("name").(string))
//}
//
//func func5() {
//	ctx := context.WithValue(context.Background(), "name", "葛诗颖")
//	go func4(ctx)
//	time.Sleep(time.Second)
//}
