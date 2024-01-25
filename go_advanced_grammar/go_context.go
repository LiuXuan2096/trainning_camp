package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

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

// 模拟一个耗时操作，比如数据库查询或者网络请求
func slowOperation(ctx context.Context) (string, error) {
	// 获取Request id
	fmt.Println("get request_id from ctx", ctx.Value("request_id"))

	// 使用一个 select 语句来监听 Context 的状态变化
	select {
	case <-time.After(3 * time.Second): // 模拟操作需要3s才能完成
		return "Done", nil
	case <-ctx.Done(): // 如果Context 被取消或者超时，返回相应的错误信息
		return "", ctx.Err()
	}
}

// 模拟一个HTTP服务器的处理函数，使用Context来控制超时
func handler(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	fmt.Println("gen request_id ", requestID)
	// 传递request_id
	ctr := context.WithValue(r.Context(), "request_id", requestID)

	// 从请求中获取 Context 并设置一个 1 秒钟的超时时间
	ctx, cancel := context.WithTimeout(ctr, 1*time.Second)

	// 在函数返回时调用取消函数
	defer cancel()

	// 调用耗时操作，并传递Context
	result, err := slowOperation(ctx)
	if err != nil {
		// 如果操作失败，返回错误信息和状态码
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	// 如果操作成功，返回结果
	fmt.Fprintln(w, result)
}

//func main() {
//	// 创建一个HTTP服务器，并注册处理函数
//	http.HandleFunc("/", handler)
//	fmt.Println("Server is running on http://localhost:9090")
//	// 启动服务器，并监听9090端口
//	err := http.ListenAndServe(":9090", nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
