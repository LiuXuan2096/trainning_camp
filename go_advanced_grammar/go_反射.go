package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	// fun1()
//	// func2()
//	// func3()
//	// func4()
//	// func5()
//	// func6()
//	// func7()
//	// func8()
//	// func9()
//	// func10()
//	// func11()
//	// func12()
//	func13()
//}
//
//type Student struct {
//	Name  string
//	Age   int
//	Score float64
//}
//
//func fun1() {
//	var num int64 = 100
//	t1 := reflect.TypeOf(num)
//	fmt.Println(t1.String())
//
//	s := Student{
//		Name: "葛诗颖",
//		Age:  22,
//	}
//	t2 := reflect.TypeOf(s)
//	fmt.Println(t2.String())
//}
//
//func func2() {
//	var num int64 = 100
//	v1 := reflect.ValueOf(num)
//	fmt.Println(v1)
//	fmt.Println(v1.String())
//
//	st := Student{
//		Name: "葛诗颖",
//		Age:  18,
//	}
//	v2 := reflect.ValueOf(st)
//	fmt.Println(v2)
//	fmt.Println(v2.String())
//}
//
//type WrapInt int
//
//func func3() {
//	var num1 int = 100
//	var num2 WrapInt = 1000
//
//	num1 = int(num2) // 不同类型的type赋值，这里要强制转换
//
//	typeNum1 := reflect.TypeOf(num1)
//	fmt.Printf("type of num1 is %s\n", typeNum1.String())
//
//	typeNum2 := reflect.TypeOf(num2)
//	fmt.Printf("type of num2 is %s\n", typeNum2.String())
//
//	fmt.Printf("kind of num1 is %v\n", typeNum1.Kind())
//	fmt.Printf("kind of num2 is %v\n", typeNum2.Kind())
//}
//
//func func4() {
//	st := Student{
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 95.5,
//	}
//	v := reflect.ValueOf(st)
//	fmt.Printf("the field num of Student is %d\n", v.NumField())
//	fmt.Printf("field1 type is %v, value is %s\n", v.Field(0).Type().Name(), v.Field(0).String())
//	fmt.Printf("field2 type is %v, value is %d\n", v.Field(1).Type().Name(), v.Field(1).Int())
//	fmt.Printf("field2 type is %v, value is %f\n", v.Field(2).Type().Name(), v.Field(2).Float())
//}
//
//func func5() {
//	m := map[int]uint32{
//		1: 100,
//		2: 200,
//	}
//	v := reflect.ValueOf(m)
//	for _, k := range v.MapKeys() {
//		filed := v.MapIndex(k)
//		fmt.Printf("key type is %v, key = %d; value type is %v, value = %d\n", k.Type().Name(), k.Int(), filed.Type().Name(), filed.Uint())
//	}
//
//}
//
//func func6() {
//	slice := []int{1, 2, 3}
//	v1 := reflect.ValueOf(slice)
//	for i := 0; i < v1.Len(); i++ {
//		elem := v1.Index(i)
//		fmt.Printf("%v ", elem.Interface())
//	}
//
//	fmt.Println()
//
//	nums := [3]int{4, 5, 6}
//	v2 := reflect.ValueOf(nums)
//	for i := 0; i < v2.Len(); i++ {
//		elem := v2.Index(i)
//		fmt.Printf("%v ", elem.Interface())
//	}
//}
//
//func func7() {
//	st := &Student{
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 98,
//	}
//	t := reflect.TypeOf(st)
//
//	fmt.Println(t.Kind())
//	fmt.Println(t.Elem().Name()) // 这里一定要加Elem()根据指针获得具体类型后，才能获得具体的type
//	fmt.Println(t.Elem().NumField())
//	for i := 0; i < t.Elem().NumField(); i++ {
//		fmt.Printf("field name is %s, field1 type is %s\n", t.Elem().Field(i).Name, t.Elem().Field(i).Type.String())
//	}
//}
//
//func Add(num1, num2 int) (int, error) {
//	return num1 + num2, nil
//}
//
//func func8() {
//	fmt.Println("input:")
//	t := reflect.TypeOf(Add)
//	for i := 0; i < t.NumIn(); i++ {
//		tIn := t.In(i)
//		fmt.Print(tIn.Name())
//		fmt.Printf(" ")
//	}
//	fmt.Printf("\n-------------------------------------")
//
//	fmt.Println("output:")
//	for i := 0; i < t.NumOut(); i++ {
//		tOut := t.Out(i)
//		fmt.Print(tOut.Name())
//		fmt.Print(" ")
//	}
//}
//
//func (s *Student) GetName() string {
//	return s.Name
//}
//
//func (s *Student) SetName(name string) {
//	s.Name = name
//}
//
//func (s *Student) SetAge(age int) {
//	s.Age = age
//}
//
//func (s *Student) GetAge() int {
//	return s.Age
//}
//
//func (s *Student) GetScore() float64 {
//	return s.Score
//}
//
//func (s *Student) SetScore(score float64) {
//	s.Score = score
//}
//
//func func9() {
//	st := &Student{
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 98,
//	}
//	t := reflect.TypeOf(st)
//
//	for i := 0; i < t.NumMethod(); i++ {
//		m := t.Method(i)
//		fmt.Printf("%+v\n", m)
//	}
//}
//
//func func10() {
//	st := &Student{
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 98,
//	}
//	fmt.Printf("st === %+v\n", st)
//
//	t := reflect.TypeOf(st)
//	v := reflect.ValueOf(st)
//
//	m1, ok := t.MethodByName("SetName") // 获取setName方法
//	fmt.Printf("t get func by name:%t\n", ok)
//
//	argsV1 := make([]reflect.Value, 0)
//	argsV1 = append(argsV1, v)
//	argsV1 = append(argsV1, reflect.ValueOf("诗诗"))
//	m1.Func.Call(argsV1)
//	fmt.Printf("st === %+v\n", st)
//
//	m2 := v.MethodByName("SetName")
//	argsV2 := make([]reflect.Value, 0)
//	argsV2 = append(argsV2, reflect.ValueOf("诗颖"))
//	m2.Call(argsV2)
//	fmt.Printf("st === %+v\n", st)
//}
//
//func func11() {
//	st := &Student{
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 98,
//	}
//	v := reflect.ValueOf(st)
//	fmt.Println(v.CanAddr())
//}
//
//func func12() {
//	st := &Student{
//		Name:  "葛诗颖",
//		Age:   22,
//		Score: 98,
//	}
//	v := reflect.ValueOf(st)
//	fmt.Println(v.Elem().CanAddr())
//}
//
//func func13() {
//	slice := []int{1, 2, 3, 4, 5}
//	v := reflect.ValueOf(slice)
//	fmt.Println(v.Index(0).CanAddr())
//	fmt.Println(v.Index(1).CanAddr())
//}
