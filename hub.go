package main

import "github.com/sdttttt/go-tds/provider"

// Multiple addresses for one service
type servicesAddress = [](*provider.Address)

// ServiceHub .
type ServiceHub struct {
	providers map[string]servicesAddress
}

// Join is Join to the service center
func (hub *ServiceHub) Join(serviceName string, service *provider.Address) {
	hub.providers[serviceName] = append(hub.providers[serviceName], service)
}

// ServiceInfo is a Service Infomation.
func (hub *ServiceHub) ServiceInfo(name string) *provider.Address {
	// TODO: constants
	return hub.providers[name][0]
}

// Start is Get ServiceHub Instance.
func Start() *ServiceHub {

	return &ServiceHub{
		// TODO: constants
		make(map[string]servicesAddress, 24),
	}
}
