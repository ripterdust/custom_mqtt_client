package main

import "github.com/ripterdust/custom_mqtt_client.git/pkg/server"

func main(){
  srv := server.NewServer()

  srv.Listen("8080")
}
