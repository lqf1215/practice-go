package main

import (
	"fmt"
	"rabbitMQ-go/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitMQ := RabbitMQ.NewRabbitMQPubSub("" + "newProduct")
	for i := 0; i < 100; i++ {
		rabbitMQ.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		fmt.Println("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}
}
