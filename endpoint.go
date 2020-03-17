package main

import "github.com/sdttttt/go-tds/provider"

// EndPoint is Get Service Entrance.
type EndPoint struct {
	hub *ServiceHub
}

// GetService is get Service Info RPC service
func (ep *EndPoint) GetService(serviceName string, service *provider.Service) error {
	service = ep.hub.ServiceInfo(serviceName)
	return nil
}
