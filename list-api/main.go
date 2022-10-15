package main

import (
	"log"
	"net/http"

	"api/message"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	router := gin.Default()
	router.GET("/list/:from/:to", func(c *gin.Context) {
		fromName := c.Param("from")
		toName := c.Param("to")
		messages := getAllMessages(redisClient, fromName, toName)
		log.Printf("Messages: %v", messages)
		c.JSON(http.StatusOK, messages)
	})

	router.Run("localhost:8081")
}

func getAllMessages(redisClient *redis.Client, fromName, toName string) []message.Message {
	var allmsgs []message.Message
	var messages []message.Message
	err := redisClient.LRange("messages", 0, -1).ScanSlice(&allmsgs)
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}

	for _, msg := range allmsgs {
		if msg.Sender == fromName && msg.Receiver == toName {
			messages = append(messages, msg)
		}
	}
	return messages
}
