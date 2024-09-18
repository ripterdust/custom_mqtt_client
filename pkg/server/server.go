package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
  r.GET("/topic", s.handleWebSocket)
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *HttpServer) handleWebSocket(c *gin.Context) {
  fmt.Println("Connecting")
  conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

  if err != nil {
   c.JSON(500, gin.H{"error": err.Error()})
   return
  }
  
  defer conn.Close()
  
  fmt.Println("Connected")
  exists, queue := s.broker.Get("queue-1")
  
  if !exists {
    errorMessage := "Queue does not exist"
    err := conn.WriteMessage(websocket.TextMessage, []byte(errorMessage))
    
    if err != nil {
      fmt.Println("Error al enviar mensaje:", err)
    }
    conn.Close()

    return
  }
  for {
   conn.WriteJSON(queue.GetAll())
   time.Sleep(time.Second)
  }

}
