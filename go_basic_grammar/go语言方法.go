package main

//import "fmt"
//
//type Student struct {
//	ID    int
//	Score int
//	People
//}
//
//type People struct {
//	Name string
//	Age  int
//}
//
//func (st *Student) GetName() string {
//	return st.Name
//}
//
//func main() {
//	// func1()
//	// func2()
//	funn3()
//}
//
///*
//在go语言中，类型的定义和绑定在类型上的方法定义可以不放在同一个文件中，但必须在同一个包下面
//*/
//func func1() {
//	st := &Student{
//		ID: 100,
//		// Name:  "葛诗颖",
//		// Age:   18,
//		Score: 98,
//	}
//	fmt.Printf("女生的姓名是: %s\n", st.GetName())
//}
//
//func (st *Student) SetScore(score int) {
//	st.Score = score
//}
//
//func (st Student) GetScore() int {
//	return st.Score
//}
//
//func func2() {
//	st := &Student{
//		ID: 100,
//		// Name:  "葛诗颖",
//		// Age:   18,
//		Score: 98,
//	}
//	fmt.Printf("设置前，女生的分数是: %d\n", st.GetScore())
//	st.SetScore(100)
//	fmt.Printf("设置后，女生的分数是: %d\n", st.GetScore())
//}
//
//func funn3() {
//	st := &Student{
//		ID:    100,
//		Score: 98,
//		People: People{
//			Name: "葛诗颖",
//			Age:  22,
//		},
//	}
//	fmt.Printf("女生的分数是: %d\n", st.GetScore())
//	fmt.Printf("女生的姓名是: %s\n", st.GetName())
//}
