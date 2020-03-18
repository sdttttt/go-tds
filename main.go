package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/sdttttt/go-tds/configuration"
)

func main() {
	hub := Start()

	receiver := &Receiver{hub}
	endpoint := &EndPoint{hub}

	// RPC interface for service provider
	rpc.Register(receiver)

	// RPC interface for service customer
	rpc.Register(endpoint)

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener Failed")
	}

	println("Version: " + configuration.Version)

	go http.Serve(listener, nil)

	select {}
}
