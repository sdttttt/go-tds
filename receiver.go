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
func (recv *Receiver) JoinServiceHub(in client.ProviderInfo, out *bool) error {

	addr := &provider.Address{IP: in.Name, Port: in.Port}

	recv.hub.Join(in.Name, addr)

	*out = true
	return nil
}
