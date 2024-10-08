package broker

import (
	"fmt"
	"sync"
  "math/rand"
  "time"

	"github.com/ripterdust/custom_mqtt_client.git/pkg/queue"
)

type Broker struct {
  queues map[string]*queue.Queue
  lock sync.Mutex
}

func NewBroker() *Broker {
  return &Broker{
    queues: make(map[string]*queue.Queue),
  }
}

func (b *Broker) ProcessQueue(name string) {
  queue := b.queues[name]
  for {

    if queue.IsEmpty() {
      timer := rand.Intn(10)
      
      time.Sleep(time.Duration(timer) * time.Second)
      
      continue
    }
    timer := rand.Intn(10)
    time.Sleep(time.Duration(timer) * time.Second)
  
    queue.Deque()
  }
}


func (b *Broker) createQueueIfNotExists(name string){
  if _,exists := b.queues[name]; exists {
    return
  }

  b.queues[name] = &queue.Queue{}
  go b.ProcessQueue(name)

}


func (b *Broker) Publish(queueName string, msg string) (bool, string) {
  b.createQueueIfNotExists(queueName)
  queue := b.queues[queueName]
  message := queue.CreateMessage(msg)
  queue.Enqueue(message)  
  
  b.queues[queueName] = queue

  return false, "Message published successfully"
}

func (b *Broker) Get(queueName string) (bool, *queue.Queue) {
  q, exists := b.queues[queueName]

  if !exists {
        return false, &queue.Queue{}
  }  
  
  return true, q
}

func (b *Broker) Topics() []string {
  keys := []string{}

  for key := range b.queues {
    keys = append(keys, key)
  }

  return keys
}
