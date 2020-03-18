package main

import (
	"github.com/sdttttt/go-tds/provider"
	"github.com/sdttttt/go-tds/trpc"
)

// EndPoint is Get Service Infomation Entrance.
// For Service customer.
type EndPoint struct {
	hub *ServiceHub
}

// GetService is get Service Info RPC service
func (ep *EndPoint) GetService(in *trpc.CustomerInfo, out *provider.Address) error {
	serviceName := in.ServiceName

	out = ep.hub.ServiceInfo(serviceName)

	return nil
}
