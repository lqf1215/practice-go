package main

import "rabbitMQ-go/RabbitMQ"

func main() {
	//Routing模式 路由模式
	//kutengOne := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	//kutengOne.RecieveRouting()

	//topic 模式 主题模式
	kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.*.one")
	kutengOne.RecieveTopic()
}
