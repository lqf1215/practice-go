package main

import (
	"fmt"
	"rabbitMQ-go/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.PublishSimple("Hello kuteng222!")
	fmt.Println("发送成功")
}
