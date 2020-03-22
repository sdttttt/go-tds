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

	addr := recv.toAddress(info)
	recv.hub.Join(info.ServiceName, addr)

	// Result: true is Successful.
	return &proto.JoinResult{Result: true}, nil
}

// ReportActive is Receive the heartbeat packet of the service.
func (recv *Receiver) ReportActive(ctx context.Context, info *proto.ProviderInfo) (*proto.ReportResult, error) {
	addr := recv.toAddress(info)
	recv.hub.isKeepLive(info.ServiceName, addr)

	return &proto.ReportResult{Result: true}, nil
}

func (recv *Receiver) toAddress(info *proto.ProviderInfo) *Address {
	return &Address{IP: info.Ip, Port: info.Port}
}
