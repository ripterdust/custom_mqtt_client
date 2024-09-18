package queue

import "github.com/gin-gonic/gin"

type Message struct {
  Id          string
  Content     string
}

type Queue struct {
  messages []Message
}

func (q *Queue) Enqueue(msg Message){

  q.messages = append(q.messages, msg)
}


func (q *Queue) RouteEnqueu(g *gin.Context){}


