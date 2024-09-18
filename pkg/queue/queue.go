package queue

type Message struct {

}

type Queue struct {
  messages []Message
}

func (q *Queue) Enqueue(msg Message){

  q.messages = append(q.messages, msg)
}


