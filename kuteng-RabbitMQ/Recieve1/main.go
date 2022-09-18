package main

import "rabbitMQ-go/RabbitMQ"

func main() {
	kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengone.RecieveRouting()
}
