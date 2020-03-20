package main

// Multiple addresses for one service
type servicesAddress = []*Address

// ServiceHub .
type ServiceHub struct {
	providers map[string]servicesAddress
}

// Address is Service Addr
type Address struct {
	IP   string
	Port string
}

// Join is Join to the service center
func (hub *ServiceHub) Join(serviceName string, service *Address) {
	hub.providers[serviceName] = append(hub.providers[serviceName], service)
}

// ServiceInfo is a Service Infomation.
func (hub *ServiceHub) ServiceInfo(name string) *Address {
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
