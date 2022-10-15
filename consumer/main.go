package main

import (
	"encoding/json"
	"log"

	"api/message"

	"github.com/go-redis/redis"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var redisClient *redis.Client

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	conn, err := amqp.Dial("amqp://user:password@localhost:7001/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MessageQueue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var m message.Message
			err := json.Unmarshal(d.Body, &m)
			if err != nil {
				log.Printf("Failed to unmarshall byte array to message: %v", err.Error())
			}

			err = saveRedis(m)
			if err != nil {
				log.Printf("Failed to save to Redis: %v", err.Error())
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func saveRedis(m message.Message) error {
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	redisClient.Do("RPUSH", "messages", bytes)

	log.Printf("Pushed message" + string(bytes))
	return nil
}
