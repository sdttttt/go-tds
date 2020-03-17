package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	hub := Start()

	receiver := &Receiver{hub: hub}

	rpc.Register(receiver)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener Failed")
	}

	http.Serve(listener, nil)
}
