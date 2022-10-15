package internal

import (
	"api/message"
	"api/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// PostMessage handler for get message from user to push the rabbitmq queue
func PostMessage(c *gin.Context) {
	var req message.Message

	// Bind the request to message type
	if err := c.BindJSON(&req); err != nil {
		return
	}

	if !req.Validate() {
		c.IndentedJSON(http.StatusBadRequest, nil)
	} else {
		err := util.PushMessage(req)
		if err != nil {
			log.Printf("Error pushing message to rabbitmq: %v", err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, req)
	}
}
