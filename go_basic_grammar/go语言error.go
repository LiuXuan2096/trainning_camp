package main

//import (
//	"errors"
//	"fmt"
//)
//
//func getPositiveSelfAdd(num int) (int, error) {
//	if num <= 0 {
//		return -1, fmt.Errorf("num is not a positive number")
//	}
//	return num + 1, nil
//}
//
//func func1() {
//	num1, err1 := getPositiveSelfAdd(1)
//	fmt.Printf("nums is %d, err is %v\n", num1, err1)
//
//	num2, err2 := getPositiveSelfAdd(-2)
//	fmt.Printf("nums is %d, err is %v\n", num2, err2)
//
//	err3 := errors.New("葛诗颖")
//	err4 := errors.New("葛诗颖")
//	fmt.Println(err3 == err4)
//
//	fmt.Println(err3.Error() == err4.Error())
//}
//
//func main() {
//	// func1()
//	func2()
//}
//
//type MyError struct {
//	code int
//	msg  string
//}
//
//func (m MyError) Error() string {
//	return fmt.Sprintf("code:%d, msg:%v", m.code, m.msg)
//}
//
//func newError(code int, msg string) error {
//	return MyError{
//		code: code,
//		msg:  msg,
//	}
//}
//
//func Code(err error) int {
//	if e, ok := err.(MyError); ok {
//		return e.code
//	}
//	return -1
//}
//
//func Msg(err error) string {
//	if e, ok := err.(MyError); ok {
//		return e.msg
//	}
//	return ""
//}
//
//func func2() {
//	err := newError(100, "test MyError")
//	fmt.Printf("code is %d, msg is %s", Code(err), Msg(err))
//}
