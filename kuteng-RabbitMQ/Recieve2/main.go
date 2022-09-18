package main

import "rabbitMQ-go/RabbitMQ"

func main() {

	// Routing模式 路由模式
	//kutengTwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	//kutengTwo.RecieveRouting()

	//topic 模式 主题模式
	kutengTwo := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.*.two")
	kutengTwo.RecieveTopic()
}
