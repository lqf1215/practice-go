package main

import (
	"fmt"
	"rabbitMQ-go/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	//Routing 模式 路由模式
	//ketengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
	//ketengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
	//for i := 0; i < 100; i++ {
	//	ketengone.PublishRouting("Hello kuteng one!" + strconv.Itoa(i))
	//	ketengtwo.PublishRouting("Hello kuteng two!" + strconv.Itoa(i))
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(i)
	//}

	//	Topic模式 主题模式

	kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.one")
	kutengTwo := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.two")
	for i := 0; i <= 100; i++ {
		kutengOne.PublishTopic("Hello kuteng topic one!" + strconv.Itoa(i))
		kutengTwo.PublishTopic("Hello kuteng topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
