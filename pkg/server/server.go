package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ripterdust/custom_mqtt_client.git/pkg/broker"
)

type HttpServer struct {
  port string
  broker *broker.Broker
}

func NewServer(broker *broker.Broker) (*HttpServer) {
  return &HttpServer {
    port: ":8080",
    broker: broker,
  }
}

func (s *HttpServer) StartServer() {
  r := gin.Default()

  r.POST("/send", s.handleSendMessage)

  r.Run()
}

func (s *HttpServer) handleSendMessage(c *gin.Context) {
  var message sendMessageRequest
  
  if err := c.BindJSON(&message); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "ok": false,})

    return
  }

  err, msg := s.broker.Publish(message.Name, message.Content)
  
  c.JSON(http.StatusOK, gin.H{
    "message": msg,
    "ok": !err,
  })
}

