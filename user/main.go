package main

import (
	"github.com/micro/user/handlers"
	"log"
	"net"
	"net/http"
)

func main() {

	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	httpHandler := handlers.NewHandler()
	s := &http.Server{
		Handler: httpHandler,
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	} else {
		log.Printf("Listening on port: %s", addr)
	}

}
