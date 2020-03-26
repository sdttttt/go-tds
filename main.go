package main

import (
	"log"
	"net"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/proto"
	"google.golang.org/grpc"
)

// Options is Hub Config Options.
type Options struct {
	balance Balance
}

// NewOptions return new Option instance.
func NewOptions(conf *configuration.Config) *Options {
	return &Options{
		balance: balanceFunc(conf.Hub.LoadBalance),
	}
}

func main() {
	config := configuration.GetConfig()

	options := NewOptions(config)

	hub := Start(options)

	server := grpc.NewServer()

	proto.RegisterReceiverServer(server, &Receiver{hub})
	proto.RegisterEndPointServer(server, &EndPoint{hub})

	listener, err := net.Listen("tcp", ":"+config.Hub.Port)
	if err != nil {
		log.Fatal("listener Failed")
	}

	println("Version: " + configuration.Version)

	server.Serve(listener)
}
