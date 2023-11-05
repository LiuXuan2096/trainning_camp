package main

//import "fmt"
//
//func main() {
//	// func1()
//	// func2()
//	// func3()
//	// func4()
//	// func5()
//	// func6()
//	func7()
//}
//
//type Phone interface {
//	Call()
//	SendMessage()
//}
//
//type Apple struct {
//	PhoneName string
//}
//
//func (a Apple) Call() {
//	fmt.Printf("%s有打电话功能\n", a.PhoneName)
//}
//
//func (a Apple) SendMessage() {
//	fmt.Printf("%s有发短信功能\n", a.PhoneName)
//}
//
//type HuaWei struct {
//	PhoneName string
//}
//
//func (h HuaWei) Call() {
//	fmt.Printf("%s有打电话功能\n", h.PhoneName)
//}
//
//func (h HuaWei) SendMessage() {
//	fmt.Printf("%s有发短信功能\n", h.PhoneName)
//}
//
//func func1() {
//	a := Apple{"Iphone"}
//	b := HuaWei{"Mate60"}
//	a.Call()
//	a.SendMessage()
//	b.Call()
//	b.SendMessage()
//
//	var phone Phone
//	phone = new(Apple)
//	phone.(*Apple).PhoneName = "iPhone"
//	phone.Call()
//	phone.SendMessage()
//}
//
//type MyWriter interface {
//	MyWriter(note string)
//}
//
//type MyReader interface {
//	MyReader()
//}
//
//type ReadWriter struct {
//}
//
//func (wr ReadWriter) MyWriter(note string) {
//	fmt.Println("Call ReadWriter MyWriter, Note =", note)
//}
//
//func (wr ReadWriter) MyReader() {
//	fmt.Println("Call ReadWriter MyReader")
//}
//
//func func2() {
//	var readWriter ReadWriter
//	readWriter.MyWriter("葛诗颖")
//	readWriter.MyReader()
//}
//
//func func3() {
//	var any interface{}
//	any = 10
//	fmt.Println(any)
//
//	any = "葛诗颖"
//	fmt.Println(any)
//
//	any = true
//	fmt.Println(any)
//}
//
//func func4() {
//	var x interface{}
//	x = "葛诗颖"
//	val, ok := x.(string)
//	fmt.Printf("val is %s, ok is %t\n", val, ok)
//}
//
//func func5() {
//	var x interface{}
//	x = "golang"
//	val := x.(int)
//	fmt.Println(val)
//}
//
//type Reader interface {
//	Read() int
//}
//
//type MyReader2 struct {
//	a, b int
//}
//
//func (m *MyReader2) Read() int {
//	return m.b + m.a
//
//}
//
//func DoJob(r Reader) {
//	fmt.Printf("myReader is %d\n", r.Read())
//}
//
//func func6() {
//	myReader := &MyReader2{2, 5}
//	DoJob(myReader)
//}
//
///*
//接口嵌套就是一个接口中包含了其他接口，如果要实现外部接口，则需要实现内部嵌套的接口对颖的所有方法
//*/
//
//type A interface {
//	run1()
//}
//
//type B interface {
//	run2()
//}
//
//// 定义嵌套接口C
//type C interface {
//	A
//	B
//	run3()
//}
//
//type Runner struct {
//}
//
//func (r Runner) run1() {
//	fmt.Println("run1!!!!")
//}
//
//func (r Runner) run2() {
//	fmt.Println("run2!!!")
//}
//
//func (r Runner) run3() {
//	fmt.Println("run3!!!")
//}
//
//func func7() {
//	var runner C
//	runner = new(Runner)
//	runner.run1()
//	runner.run2()
//	runner.run3()
//}
