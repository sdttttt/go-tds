package main

import "github.com/sdttttt/go-tds/provider"

// ServiceHub .
type ServiceHub struct {
	providers map[string]*provider.Service
}

// Join is Join to the service center
func (hub *ServiceHub) Join(name string, service *provider.Service) {
	hub.providers[name] = service
}

// ServiceInfo is a Service Infomation.
func (hub *ServiceHub) ServiceInfo(name string) *provider.Service {
	return hub.providers[name]
}

// Start is Get ServiceHub Instance.
func Start() *ServiceHub {
	return &ServiceHub{
		// TODO: constants
		make(map[string]*provider.Service, 24),
	}
}
