package main

import (
	"github.com/sdttttt/go-tds/client"
	"github.com/sdttttt/go-tds/provider"
)

// Receiver is Rpc Service Entrance.
// For Service Provider
type Receiver struct {
	hub *ServiceHub
}

// JoinServiceHub is External registration service
// info is Service Info
// result is Whether the service registration is Successful.
func (recv *Receiver) JoinServiceHub(info client.ProviderInfo, result *bool) error {

	addr := &provider.Address{IP: info.Name, Port: info.Port}
	service := &provider.Service{Address: addr}

	recv.hub.Join(info.Name, service)

	*result = true
	return nil
}
