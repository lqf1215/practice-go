package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

// rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	// 交换机名称
	Exchange string
	// bind key 名称
	Key string
	// 连接信息
	Mqurl string
}

//创建结构体实列
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}

}

// 断开channel 和connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 创建简单模式RabbitMQ 实列
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitm := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitm.conn, err = amqp.Dial(rabbitm.Mqurl)
	rabbitm.failOnErr(err, "failed to connect rabbitmq")
	// 获取channel
	rabbitm.channel, err = rabbitm.conn.Channel()
	rabbitm.failOnErr(err, "failed to open a channel")
	return rabbitm
}

// 直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {

	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil)

	if err != nil {
		fmt.Println(err)
	}

	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	// 接收信息
	msgs, err := r.channel.Consume(q.Name,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否独有
		false,
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,
		//列是否阻塞
		false,
		nil)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for msg := range msgs {
			log.Printf("Received a message:%s", msg.Body)
		}
	}()
	fmt.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// 订阅模式创建Rabbit MQ实列
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	// 创建Rabbit MQ实列
	rabbit := NewRabbitMQ("", exchangeName, "")
	var err error
	rabbit.conn, err = amqp.Dial(rabbit.Mqurl)
	rabbit.failOnErr(err, "failed to connect rabbitmq!")
	// 获取channel
	rabbit.channel, err = rabbit.conn.Channel()
	rabbit.failOnErr(err, "failed to open a channel")
	return rabbit
}

// 订阅模式生产
func (r *RabbitMQ) PublishPub(msg string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true, false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false, false, nil)
	r.failOnErr(err, "Failed to declare an exchannge")

	err = r.channel.Publish(
		r.Exchange,
		"",
		false, false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(msg)})

}

//订阅模式消费端代码
func (r *RabbitMQ) RecieveSub() {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exch"+
		"ange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range messges {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("退出请按 CTRL+C\n")
	<-forever
}
