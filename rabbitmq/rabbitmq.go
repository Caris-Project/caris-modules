package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/streadway/amqp"
)

var (
	connection *amqp.Connection
	channel    *amqp.Channel
)

// Init ...
func Init(uri string) *amqp.Connection {
	conn, err := amqp.Dial(uri)
	if err != nil {
		fmt.Println(aurora.Red("*** Connect to RabbitMQ failed: " + uri))
		panic(err)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO RABBITMQ: " + uri))

	// Set connection
	connection = conn

	// Set channel
	channel, _ = connection.Channel()

	return connection
}

// Log ...
func Log(queueName string) {
	fmt.Println(aurora.BgBlue("*** New message from queue: " + queueName))
}

// PublishMessage ...
func PublishMessage(queueName string, data interface{}) error {
	byteData, _ := json.Marshal(data)

	// Set publishing data
	message := generateData(byteData)

	// Send
	return channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		message,
	)
}

// generateData ...
func generateData(body []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
}

// PublishErrorLog ...
func PublishErrorLog(queueName string, err error) {
	fmt.Println("*** Error when publish message to", queueName, err.Error())
}

// DeclareQueue ...
func DeclareQueue(queueName string) {
	_, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		fmt.Println("... Declare failed subscriber queue "+queueName, err)
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
