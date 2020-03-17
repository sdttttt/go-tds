package main

import (
	"github.com/sdttttt/go-tds/client"
	"github.com/sdttttt/go-tds/provider"
)

// EndPoint is Get Service Infomation Entrance.
// For Service customer.
type EndPoint struct {
	hub *ServiceHub
}

// GetService is get Service Info RPC service
func (ep *EndPoint) GetService(info *client.CustomerInfo, service *provider.Service) error {
	serviceName := info.ServiceName
	service = ep.hub.ServiceInfo(serviceName)
	return nil
}
