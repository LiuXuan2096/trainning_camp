package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	////1. 创建http server对象
	//mux := http.NewServeMux()
	//
	//// 2. 注册路由和中间件
	//mux.HandleFunc("/girlFriend", func(resp http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(resp, "诗诗")
	//})
	//mux.HandleFunc("/wife", func(resp http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(resp, "葛诗颖")
	//})
	////3.启动http服务器，传入监听地址和多路复用器
	//err := http.ListenAndServe(":8888", mux)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//http.HandleFunc("/girlFriend", func(resp http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(resp, "诗诗")
	//})
	//http.HandleFunc("/wife", func(rest http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(rest, "葛诗颖")
	//})
	//
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}

	http.HandleFunc("/login", loginHandler)  // 设置访问的路由
	err := http.ListenAndServe(":8888", nil) // 设置监听的端口
	if err != nil {
		fmt.Println(err)
	}

}

func loginHandler(rep http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(rep, "err %v", err)
			return
		}
		err = t.Execute(rep, nil)
		if err != nil {
			fmt.Fprintf(rep, "err %v", err)
		}
	} else {
		fmt.Println(" req.Form: ", req.Form)      // req.Form: map[]
		fmt.Println(" req.Form", req.Form == nil) // req.Form: true
		fmt.Println("req.PostForm:", req.PostForm)
		fmt.Println("FormValue:", req.FormValue("username"))
		fmt.Println("PostFormValue", req.PostFormValue("username"))
		fmt.Println("req.PostForm:", req.PostForm)
		fmt.Println("FormValue:param:", req.FormValue("param"))
		fmt.Println("PostFormValue:param", req.PostFormValue("param"))
		fmt.Println(" req.Form:", req.Form)
	}
}
