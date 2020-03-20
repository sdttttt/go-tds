package main

import (
	"context"

	"github.com/sdttttt/go-tds/proto"
)

// Receiver is Rpc Service Entrance.
// For Service Provider
type Receiver struct {
	hub *ServiceHub
}

// JoinServiceHub is External registration service
// info is Service Info
// result is Whether the service registration is Successful.
func (recv *Receiver) JoinServiceHub(ctx context.Context, info *proto.ProviderInfo) (*proto.JoinResult, error) {

	addr := &Address{IP: info.Ip, Port: info.Port}
	recv.hub.Join(info.ServiceName, addr)

	// Result: true is Successful.
	return &proto.JoinResult{Result: true}, nil
}
