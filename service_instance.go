package main

import (
	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/utils"
)

// ServiceInstance is a basic service Information.
type ServiceInstance struct {
	ip   string
	port string

	group *ServiceGroup
	timer *utils.Timer
}

// removeServiceInstance from ServiceGroup
func removeServiceInstance(group *ServiceGroup, instance *ServiceInstance) func() {
	return func() {
		group.remove(instance)
	}
}

// NewServiceInstance is initializer a new ServiceInstance.
func NewServiceInstance(group *ServiceGroup, ip string, port string) *ServiceInstance {
	timer := utils.NewTimer()
	instance := &ServiceInstance{
		ip,
		port,
		group,
		timer,
	}

	config := configuration.GetConfig()

	// Automatic deletion of service timeout.
	instance.timer.AddJob(config.Hub.CheckSurvivalTime, destroy(instance))
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
	instance.timer = utils.NewTimer()
	config := configuration.GetConfig()
	instance.timer.AddJob(config.Hub.CheckSurvivalTime, destroy(instance))
	instance.timer.Start()
}

// destroy return func() is Remove ServiceInstance from ServiceGroup.
func destroy(instance *ServiceInstance) func() {
	group := instance.group
	return func() {
		group.remove(instance)
	}
}
