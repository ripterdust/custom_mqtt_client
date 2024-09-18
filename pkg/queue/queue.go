package queue

import "github.com/google/uuid"

type Message struct {
  Id          string
  Content     string
}

type Queue struct {
  messages []Message
}

func (q *Queue) Enqueue(message Message){
  q.messages = append(q.messages, message)
}


func (q *Queue) GetAll() []Message {
  return q.messages
}

func (q *Queue) Deque() {}

func generageId() string {
  return uuid.New().String()
}

func (q *Queue) CreateMessage(content string) Message {
  return Message {
    Content: content,
    Id: generageId(),
  }

}
