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
	if hub.providers[name] == nil {
		return false
	}
	if hub.providers[name].useLen <= 0 {
		return false
	}

	return true
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

// isKeepLive is The service is Can be use.
func (hub *ServiceHub) isKeepLive(serviceName string, addr *Address) {
	hub.providers[serviceName].serviceIsLive(addr)
}

// Start is Get ServiceHub Instance.
func Start() *ServiceHub {
	return &ServiceHub{
		// TODO: constants
		make(map[string]*ServiceGroup, 0),
	}
}
