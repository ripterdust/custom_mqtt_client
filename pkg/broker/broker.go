package broker

import (
	"fmt"
	"sync"
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

func (b *Broker) createQueueIfNotExists(name string){
  b.lock.Lock()

  if _,exists := b.queues[name]; !exists {
    b.queues[name] = &queue.Queue{}
  }

  defer b.lock.Unlock()
}


func (b *Broker) Publish(queueName string, msg string) (bool, string) {
  b.createQueueIfNotExists(queueName)
  
  queue := b.queues[queueName]
  message := queue.CreateMessage(msg)
  
  go b.processMessage(message.Id)

  b.lock.Lock()
  
  

  defer b.lock.Unlock()
    
  return false, "Message published successfully"
}

func (b *Broker) processMessage(messageId string){
  time.Sleep(5 * time.Second)

  fmt.Println(messageId)
}


