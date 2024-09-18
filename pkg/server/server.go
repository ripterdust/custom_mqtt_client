package server

import (
	"fmt"
	"log"
	"net/http"
)

type Httpserver struct {}

func NewServer() *Httpserver {
  return &Httpserver{}
}

func (s *Httpserver) Listen(port string) {
  
  fmt.Println("Server listening on port: " + port)
  log.Fatal(http.ListenAndServe(":" + port, nil))

}
