package util

import (
	"api/message"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func PushMessage(m message.Message) error {
	conn, err := amqp.Dial("amqp://user:password@localhost:7001/")
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err.Error())
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err.Error())
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"MessageQueue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err.Error())
		return err
	}

	byteMessage, err := json.Marshal(m)
	if err != nil {
		log.Printf("Failed to marshall message to byte array: %v", err.Error())
		return err
	}

	// Create a message to publish.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteMessage,
	}

	// Attempt to publish a message to the queue.
	if err := ch.Publish(
		"",             // exchange
		"MessageQueue", // queue name
		false,          // mandatory
		false,          // immediate
		message,        // message to publish
	); err != nil {
		return err
	}
	log.Printf("[x] Sent %s\n", m.Message)
	return nil
}
