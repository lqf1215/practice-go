package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义一个接口，有一个方法
type A interface {
	Println()
}

// 定义一个接口，有两个方法
type B interface {
	Println()
	Printf() int
}

// 定义一个结构体
type A1Instance struct {
	Data string
}

// 结构体实现了Println()方法，现在它是一个 A 接口
func (a1 *A1Instance) Println() {
	fmt.Println("a1:", a1.Data)
}

// 定义一个结构体
type A2Instance struct {
	Data string
}

// 结构体实现了Println()方法，现在它是一个 A 接口
func (a2 *A2Instance) Println() {
	fmt.Println("a2:", a2.Data)
}

// 结构体实现了Printf()方法，现在它是一个 B 接口，它既是 A 又是 B 接口
func (a2 *A2Instance) Printf() int {
	fmt.Println("a2:", a2.Data)
	return 0
}

//func main() {
//	ch := make(chan int)
//	go Hu(ch)
//	fmt.Println("start hu,wait...")
//
//	v := <-ch
//	fmt.Println("receive:", v)
//
//}

func Hu(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 secord hu!!!")
	ch <- 1000
}

//func main() {
//	producer, err := rpc.DialHTTP("tcp", "localhost:8081")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	timeStamp := time.Now().Unix()
//	request := pb.OrderRequest{OrderId: "201907300001", TimeStamp: timeStamp}
//
//	var response *pb.OrderInfo
//	err = producer.Call("OrderService.GetOrderInfo", request, &response)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	fmt.Println(*response)
//
//}

//全局变量
var ticket = 10 // 100张票

var wg sync.WaitGroup
var matex sync.Mutex // 创建锁头

func main() {
	/*
	   4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	*/
	wg.Add(4)
	go saleTickets("售票口1") // g1,100
	go saleTickets("售票口2") // g2,100
	go saleTickets("售票口3") //g3,100
	go saleTickets("售票口4") //g4,100
	wg.Wait()              // main要等待。。。

	//time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	//for i:=1;i<=100;i++{
	//  fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		matex.Lock()
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket) // 1 , 0, -1 , -2
			ticket--                         //0 , -1 ,-2 , -3
		} else {
			matex.Unlock() //解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		matex.Unlock() //解锁
	}
}
