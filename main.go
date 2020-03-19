package main

import (
	"log"
	"net"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/proto"
	"google.golang.org/grpc"
)

func main() {
	hub := Start()

	server := grpc.NewServer()

	proto.RegisterReceiverServer(server, &Receiver{hub})
	proto.RegisterEndPointServer(server, &EndPoint{hub})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener Failed")
	}

	println("Version: " + configuration.Version)

	server.Serve(listener)
}
