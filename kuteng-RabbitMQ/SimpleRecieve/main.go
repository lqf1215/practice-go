package main

import "rabbitMQ-go/RabbitMQ"

func main() {
	rabbitMQ := RabbitMQ.NewRabbitMQSimple("" + "kuteng")
	rabbitMQ.ConsumeSimple()
}
