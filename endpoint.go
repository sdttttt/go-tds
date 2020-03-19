package main

import (
	"context"

	"github.com/sdttttt/go-tds/proto"
)

// EndPoint is Get Service Infomation Entrance.
// For Service customer.
type EndPoint struct {
	hub *ServiceHub
}

// GetServiceInfo is get Service Info RPC service
func (ep *EndPoint) GetServiceInfo(ctx context.Context, in *proto.ProviderInfo) (*proto.ProviderInfo, error) {
	serviceName := in.ServiceName
	addr := ep.hub.ServiceInfo(serviceName)

	in.Ip = addr.IP
	in.Port = addr.Port

	return in, nil
}
