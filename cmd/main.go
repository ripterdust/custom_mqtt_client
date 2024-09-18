package main

import "github.com/ripterdust/custom_mqtt_client.git/pkg/server"

func main(){
  server := server.NewServer()

  server.Listen("3030")
}
