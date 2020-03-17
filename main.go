package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	hub := Start()

	receiver := &Receiver{hub}
	endpoint := &EndPoint{hub}

	rpc.Register(receiver)
	rpc.Register(endpoint)

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener Failed")
	}

	println("Version: " + VERSION)

	http.Serve(listener, nil)
}
