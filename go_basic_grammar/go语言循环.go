package main

//import "fmt"
//
//func main() {
//	// func1()
//	// func2()
//	func3()
//}
//
//func func1() {
//	for i := 0; i < 5; i++ {
//		fmt.Printf("current i %d\n", i)
//	}
//	j := 0
//	for {
//		if j == 5 {
//			break
//		}
//		fmt.Printf("current j %d\n", j)
//		j++
//	}
//	// 声明数组时 可以不指定个数
//	var strAry = []string{"aa", "bb", "cc", "dd", "葛诗颖"}
//	// 切片初始化
//	var sliceAry = make([]string, 0)
//	sliceAry = strAry[1:3]
//	for i, str := range sliceAry {
//		fmt.Printf("slice i %d, str %s\n", i, str)
//	}
//	// 字典初始化
//	var dic = map[string]int{
//		"葛诗颖": 1,
//		"诗诗":  2,
//	}
//	for k, v := range dic {
//		fmt.Printf("key %s, value %d\n", k, v)
//	}
//}
//
//func func2() {
//	arr := [2]int{1, 2}
//	res := []*int{}
//	for _, v := range arr {
//		// v 每次都是同一个变量
//		fmt.Println(&v)
//		res = append(res, &v)
//	}
//	// expect: 1,2
//	// but result : 2 2
//	fmt.Println(*res[0], *res[1])
//}
//
//func func3() {
//	/*
//		在循环遍历的同时往遍历的切片追加元素，循环依然会停止。因为遍历前，这个切片的长度已经确定了，再追加元素，
//		也不会影响到这个长度。在对切片a遍历的时候，显示将a复制给一个临时变量temp，然后根据temp的长度做遍历，
//		最终在append的时候是将取到的value追加到切片a中
//	*/
//	v := []int{1, 2, 3}
//	// length := len(v)
//	for i := range v {
//		v = append(v, i)
//	}
//}
