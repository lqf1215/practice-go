package base

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//1. 读一个已经关闭的通道
func TestName1(t *testing.T) {
	channel := make(chan int, 10)
	channel <- 2
	close(channel)
	x := <-channel
	fmt.Println(x)
}

// 遍历读关闭通道
func TestName2(t *testing.T) {
	channel := make(chan int, 10)
	channel <- 2
	channel <- 3
	close(channel) //若不关闭通道，则会报死锁错误
	for num := range channel {
		fmt.Println(num)
	}
}

//2. 写一个已经关闭的通道
func TestName3(t *testing.T) {
	channel := make(chan int, 10)
	close(channel)
	channel <- 1
}

//3. 关闭一个已经关闭的管道
func TestName4(t *testing.T) {
	channel := make(chan int, 10)
	close(channel)
	close(channel)

}

var wg sync.WaitGroup

func TestName5(t *testing.T) {

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)
	ch1 <- struct{}{}
	wg.Add(3)
	start := time.Now().Unix()
	go print("gorouine1", ch1, ch2)
	go print("gorouine2", ch2, ch3)
	go print("gorouine3", ch3, ch1)
	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("duration:%d", end-start)

}
func print(gorouine string, inputchan chan struct{}, outchan chan struct{}) {
	// 模拟内部操作耗时
	time.Sleep(1 * time.Second)
	select {
	case <-inputchan:
		fmt.Printf("%s", gorouine)
		outchan <- struct{}{}
	}
	wg.Done()
}
