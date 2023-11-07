package main

//import (
//	"fmt"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//func main() {
//	// 创建任务池
//	pool := NewPool(3, 10)
//
//	for i := 0; i < 20; i++ {
//		// 任务放入池中
//		pool.AddTask(NewTask(func() error {
//			fmt.Printf("I am a Task:\n葛诗颖\n")
//			return nil
//		}))
//	}
//
//	time.Sleep(1e9)
//}
//
//type Task struct {
//	f func() error // 具体的任务逻辑
//}
//
//func NewTask(funcArg func() error) *Task {
//	return &Task{
//		f: funcArg,
//	}
//}
//
//type Pool struct {
//	RunningWorkers int64      // 运行着的worker数量
//	Capacity       int64      // 协程池worker容量
//	JobCh          chan *Task // 用于worker取任务
//	sync.Mutex
//}
//
//func NewPool(capacity int64, taskNum int) *Pool {
//	return &Pool{
//		Capacity: capacity,
//		JobCh:    make(chan *Task, taskNum),
//	}
//}
//
//func (p *Pool) GetCap() int64 {
//	return p.Capacity
//}
//
//func (p *Pool) incRunning() {
//	atomic.AddInt64(&p.RunningWorkers, 1)
//}
//
//func (p *Pool) decRunning() {
//	atomic.AddInt64(&p.RunningWorkers, -1)
//}
//
//func (p *Pool) GetRunningWorkers() int64 {
//	return atomic.LoadInt64(&p.RunningWorkers)
//}
//
//func (p *Pool) run() {
//	p.incRunning()
//	go func() {
//		defer func() {
//			p.decRunning()
//		}()
//		for task := range p.JobCh {
//			task.f()
//		}
//	}()
//}
//
//// AddTask 往协程池里添加任务
//func (p *Pool) AddTask(task *Task) {
//	// 加锁防止启动多个worker
//	p.Lock()
//	defer p.Unlock()
//
//	if p.GetRunningWorkers() < p.GetCap() {
//		p.run()
//	}
//
//	// 将任务推入队列，等待消费
//	p.JobCh <- task
//}
