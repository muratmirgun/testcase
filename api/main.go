package main

import (
	"api/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/message", internal.PostMessage)

	router.Run("localhost:8080")
}
