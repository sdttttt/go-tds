package main

// ServiceHub .
type ServiceHub struct {
	providers map[string]*ServiceGroup
}

// Address is Service Addr
type Address struct {
	IP   string
	Port string
}

func (hub *ServiceHub) hasService(name string) bool {
	return hub.providers[name] != nil
}

// Join is Join to the service center
func (hub *ServiceHub) Join(serviceName string, service *Address) {
	if !hub.hasService(serviceName) {
		hub.providers[serviceName] =
			NewServiceGroup(0, RoundRobin())
	}

	hub.providers[serviceName].add(service)
}

// ServiceInfo is a Service Infomation.
func (hub *ServiceHub) ServiceInfo(name string) *Address {
	if hub.hasService(name) {
		return hub.providers[name].next()
	}

	return nil
}

// Start is Get ServiceHub Instance.
func Start() *ServiceHub {

	return &ServiceHub{
		// TODO: constants
		make(map[string]*ServiceGroup, 24),
	}
}
