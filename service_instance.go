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
	instance := &ServiceInstance{
		ip,
		port,
		group,
		timer,
	}

	// TODO: CONSTANTS
	instance.timer.AddFunc("0 2 * * * * *", destroy(instance))
	instance.timer.Start()

	return instance
}

// toAddress is from ServiceInstance Become Address.
func (instance *ServiceInstance) toAddress() *Address {
	return &Address{
		IP:   instance.ip,
		Port: instance.port,
	}
}

// resetSurvivalTime is Reset Survival time of ServiceInstance
func (instance *ServiceInstance) resetSurvivalTime() {
	instance.timer.Stop()
	instance.timer = cron.New()
	instance.timer.AddFunc("0 2 * * * * *", destroy(instance))
	instance.timer.Start()
}

// destroy return func() is Remove ServiceInstance from ServiceGroup.
func destroy(instance *ServiceInstance) func() {

	group := instance.group

	return func() {
		for index, currentInstance := range group.instances {
			if instance == currentInstance {
				group.instances =
					append(group.instances[:index], group.instances[index+1:]...)
			}
		}
	}
}
