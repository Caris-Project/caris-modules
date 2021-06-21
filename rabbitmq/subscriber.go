package rabbitmq

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/streadway/amqp"
)

// SubscribeQueue ...
func SubscribeQueue(queueName string) {
	_, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		fmt.Println("*** Declare failed subscriber queue "+queueName, err.Error())
	} else {
		fmt.Println(aurora.Green("*** RabbitMQ - Declared queue: " + queueName))
	}
}

// GetMessagesFromQueue ...
func GetMessagesFromQueue(queueName string) <-chan amqp.Delivery {
	messages, _ := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return messages
}
