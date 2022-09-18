package main

import (
	"fmt"
	"rabbitMQ-go/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	ketengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	ketengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	for i := 0; i < 100; i++ {
		ketengone.PublishRouting("Hello kuteng one!" + strconv.Itoa(i))
		ketengtwo.PublishRouting("Hello kuteng two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
