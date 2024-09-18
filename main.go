package main

import (
	"github.com/ripterdust/custom_mqtt_client.git/pkg/broker"
	"github.com/ripterdust/custom_mqtt_client.git/pkg/server"
)

func main(){
  broker := broker.NewBroker()
  server := server.NewServer(broker)

  server.StartServer()
}
