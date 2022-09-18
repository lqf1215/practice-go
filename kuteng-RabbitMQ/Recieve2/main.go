package main

import "rabbitMQ-go/RabbitMQ"

func main() {
	kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	kutengtwo.RecieveRouting()
}
