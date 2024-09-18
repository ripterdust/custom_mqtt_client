package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
  port string
}

func NewServer() (*HttpServer) {
  return &HttpServer {
    port: ":8080",
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

  fmt.Println("----------- MESSAGE ----------------")
  fmt.Println(message)

  c.JSON(http.StatusOK, gin.H{
    "message": "message published",
    "ok": true,
    "data": message,
  })
}

