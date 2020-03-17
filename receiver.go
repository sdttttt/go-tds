package main

import (
	"github.com/sdttttt/go-tds/msg"
	"github.com/sdttttt/go-tds/provider"
)

// Receiver is Rpc Service Entrance.
type Receiver struct {
	hub *ServiceHub
}

// JoinServiceHub is External registration service
// info is Service Info
// result is Whether the service registration is Successful.
func (recv *Receiver) JoinServiceHub(info msg.ProviderInfo, result *bool) error {

	addr := &provider.Address{IP: info.Name, Port: info.Port}
	service := &provider.Service{Address: addr}

	recv.hub.Join(info.Name, service)

	*result = true
	return nil
}
