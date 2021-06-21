package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// Publish new message to channel
func Publish(queueName string, data interface{}) error {
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
