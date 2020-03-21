package main

import "github.com/robfig/cron/v3"

type ServiceInstance struct {
	ip   string
	port string

	group *ServiceGroup

	timer *cron.Cron
}

// removeServiceInstance from ServiceGroup
func removeServiceInstance(group *ServiceGroup, instance *ServiceInstance) func() {
	return func() {
		group.remove(instance)
	}
}

// NewServiceInstance is initializer a new ServiceInstance.
func NewServiceInstance(group *ServiceGroup, ip string, port string) *ServiceInstance {
	timer := cron.New()
	return &ServiceInstance{
		ip,
		port,
		group,
		timer,
	}
}

func (instance *ServiceInstance) toAddress() *Address {
	return &Address{
		IP:   instance.ip,
		Port: instance.port,
	}
}
