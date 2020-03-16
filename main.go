package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	hub := Start()

	rece := &Receiver{Hub: hub}

	rpc.Register(rece)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener Faild")
	}

	http.Serve(listener, nil)
}
