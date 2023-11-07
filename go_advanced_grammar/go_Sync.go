package main

//
//import (
//	"fmt"
//	"strconv"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//// 声明配置结构体Config
//type Config struct {
//}
//
//var wg sync.WaitGroup
//var instance *Config
//var once sync.Once
//var num int
//var mu = sync.Mutex{}
//
//func main() {
//	// func1()
//	// func2()
//	// InitConfig()
//	//func3()
//	// func4()
//	// func5()
//	// func6()
//	// func7()
//	// func8()
//	// func9()
//	// func10()
//	// func11()
//	func12()
//}
//
//func func1() {
//	ch := make(chan struct{}, 10)
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			fmt.Printf("num:%d\n", i)
//			ch <- struct{}{}
//		}(i)
//	}
//
//	for i := 0; i < 10; i++ {
//		<-ch
//	}
//
//	fmt.Println("end")
//}
//
//func myGoroutine() {
//	defer wg.Done()
//	fmt.Println("myGoroutine!")
//}
//
//func func2() {
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go myGoroutine()
//	}
//	wg.Wait()
//	fmt.Println("end!!!")
//}
//
//// 获取配置结构体
//func InitConfig() *Config {
//	once.Do(func() {
//		instance = &Config{}
//	})
//	return instance
//}
//
//func add() {
//	defer wg.Done()
//	num += 1
//}
//
//func func3() {
//	var n = 10 * 10 * 10 * 10 * 100
//	wg.Add(n)
//
//	for i := 0; i < n; i++ {
//		// 启动n个goroutine去累加num
//		go add()
//	}
//
//	// 等待所有的goroutine执行完毕
//	wg.Wait()
//
//	// 不出意外的话，num应该等于n，但实际上 程序的执行结果不是这样
//	fmt.Printf("num = %d, n = %d\n", num, n)
//	fmt.Println(num == n)
//}
//
//func add1() {
//	mu.Lock()
//	defer wg.Done()
//	num += 1
//	mu.Unlock()
//}
//
//func func4() {
//	var n = 10 * 1000 * 100
//	wg.Add(n)
//
//	for i := 0; i < n; i++ {
//		// 启动n个goroutine去累加num
//		go add1()
//	}
//
//	// 等待所有goroutine执行完毕
//	wg.Wait()
//	fmt.Println(num == n)
//	fmt.Printf("num = %d, n = %d\n", num, n)
//}
//
//var cnt = 0
//
//func func5() {
//	var mr sync.RWMutex
//	for i := 1; i <= 3; i++ {
//		go write(&mr, i)
//	}
//	for i := 1; i <= 3; i++ {
//		go read(&mr, i)
//	}
//
//	time.Sleep(time.Second)
//	fmt.Println("final coutn:", cnt)
//}
//
//func read(mr *sync.RWMutex, i int) {
//	fmt.Printf("goroutine%d reader start\n", i)
//	mr.RLock()
//	fmt.Printf("goroutine%d reading count:%d\n", i, cnt)
//	time.Sleep(time.Millisecond)
//	mr.RUnlock()
//
//	fmt.Printf("goroutine%d reader over\n", i)
//}
//
//func write(mr *sync.RWMutex, i int) {
//	fmt.Printf("goroutine%d writer start\n", i)
//	mr.Lock()
//	cnt++
//	fmt.Printf("goroutine%d writting count:%d\n", i, cnt)
//	time.Sleep(time.Millisecond)
//	mr.Unlock()
//
//	fmt.Printf("goroutine%d writer over\n", i)
//}
//
//func func6() {
//	var mu sync.Mutex
//	mu.Lock()
//	defer mu.Unlock()
//	copyMutex(mu)
//}
//
//func copyMutex(mu sync.Mutex) {
//	mu.Lock()
//	defer mu.Unlock()
//	fmt.Println("ok")
//}
//
//// 循环等待导致死锁
//func func7() {
//	var mu1, mu2 sync.Mutex
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		mu1.Lock()
//		defer mu1.Unlock()
//		time.Sleep(1 * time.Second)
//
//		mu2.Lock()
//		defer mu2.Unlock()
//	}()
//
//	go func() {
//		defer wg.Done()
//		mu2.Lock()
//		defer mu2.Unlock()
//		time.Sleep(1 * time.Second)
//		mu1.Lock()
//		defer mu1.Unlock()
//	}()
//	wg.Wait()
//}
//
//// go语言内置的map并不是并发安全的，在多个goroutine同时操作map的时候，会有并发问题
//var m = make(map[string]int)
//
//// var mu sync.Mutex
//
//func getVal(key string) int {
//	return m[key]
//}
//
//func setVal(key string, value int) {
//	m[key] = value
//}
//
//func func8() {
//	wg := sync.WaitGroup{}
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go func(num int) {
//			defer wg.Done()
//			key := strconv.Itoa(num)
//			setVal(key, num)
//			fmt.Printf("key=%v, val=%v\n", key, getVal(key))
//		}(i)
//	}
//	wg.Wait()
//}
//
//func func9() {
//	wg := sync.WaitGroup{}
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go func(num int) {
//			defer func() {
//				wg.Done()
//				mu.Unlock()
//			}()
//			key := strconv.Itoa(num)
//			mu.Lock()
//			setVal(key, num)
//			fmt.Printf("key=%v, val=%v\n", key, getVal(key))
//		}(i)
//	}
//	wg.Wait()
//}
//
//func func10() {
//	var m sync.Map
//	// 1.写入
//	m.Store("name", "葛诗颖")
//	m.Store("age", 22)
//
//	// 2.读取
//	age, _ := m.Load("age")
//	fmt.Println(age.(int))
//
//	// 3.遍历
//	m.Range(func(key, value any) bool {
//		fmt.Printf("key is:%v, val is:%v\n", key, value)
//		return true
//	})
//
//	// 4.删除
//	m.Delete("age")
//	age, ok := m.Load("age")
//	fmt.Println(age, ok)
//
//	// 5.读取或写入
//	m.LoadOrStore("name", "葛诗颖")
//	name, _ := m.Load("name")
//	fmt.Println(name)
//}
//
//func func11() {
//	var sum int32 = 0
//	var wg sync.WaitGroup
//	for i := 0; i < 100*100*100; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			atomic.AddInt32(&sum, 1)
//		}()
//	}
//	wg.Wait()
//	fmt.Printf("sum is %d\n", sum)
//}
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//func func12() {
//	pool := sync.Pool{
//		New: func() interface{} {
//			return &Student{
//				Name: "葛诗颖",
//				Age:  22,
//			}
//		},
//	}
//
//	st := pool.Get().(*Student)
//	println(st.Name, st.Age)
//	fmt.Printf("addr is %p\n", st)
//
//	pool.Put(st)
//	st1 := pool.Get().(*Student)
//	println(st1.Name, st1.Age)
//	fmt.Printf("addr is %p\n", st1)
//
//}
