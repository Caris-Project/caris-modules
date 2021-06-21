package rabbitmq

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/streadway/amqp"
)

var (
	connection *amqp.Connection
	channel    *amqp.Channel
)

// Connect ...
func Connect(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		fmt.Println(aurora.Red("*** Connect to RabbitMQ failed: " + uri))
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO RABBITMQ: " + uri))

	// Set connection
	connection = conn

	// Set channel
	channel, _ = connection.Channel()

	return nil
}

// GetConnection ...
func GetConnection() *amqp.Connection {
	return connection
}

// Log ...
func log(queueName string) {
	fmt.Println(aurora.BgBlue("*** New message from queue: " + queueName))
}

// PublishErrorLog ...
func PublishErrorLog(queueName string, err error) {
	fmt.Println("*** Error when publish message to", queueName, err.Error())
}
