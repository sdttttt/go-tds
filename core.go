package main

import "github.com/sdttttt/go-tds/provider"

// ServiceHub .
type ServiceHub struct {
	Providers []provider.Service
}

func (hub *ServiceHub) Join(name string) {

	service := provider.Service{Name: name}
	hub.Providers = append(hub.Providers, service)
}

// Start ServiceHub Instance.
func Start() *ServiceHub {
	return &ServiceHub{
		// TODO: constants
		make([]provider.Service, 12),
	}
}
