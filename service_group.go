package main

// ServiceGroup is The Same Service, put in this group.
type ServiceGroup struct {
	instances []*ServiceInstance
	useLen    uint8
	index     uint8

	arithmetic Balance
}

// NewServiceGroup is Get ServiceGroup Instance.
func NewServiceGroup(len uint8, bala Balance) *ServiceGroup {
	return &ServiceGroup{
		instances: make([]*ServiceInstance, len),
		useLen:    0,
		index:     0,

		arithmetic: bala,
	}
}

// add is Join to ServiceGroup
func (group *ServiceGroup) add(addr *Address) {

	instance := NewServiceInstance(group, addr.IP, addr.Port)

	group.instances = append(group.instances, instance)
	group.useLen++
}

func (group *ServiceGroup) remove(in *ServiceInstance) {
	for index, currentInstance := range group.instances {
		if currentInstance == in {
			group.instances = append(group.instances[:index], group.instances[index+1:]...)
		}
	}
}

// next is get next same Service.
func (group *ServiceGroup) next() *Address {
	if group.useLen == 0 {
		return nil
	}

	if group.index == group.useLen {
		group.index = 0
	}

	index := group.arithmetic(&group.useLen, &group.index)
	return group.instances[index].toAddress()
}

// forEach is foreach all ServiceInstance.
func (group *ServiceGroup) forEach(fn func(int, *ServiceInstance)) {
	for index, instance := range group.instances {
		fn(index, instance)
	}
}

// serviceIsLive is reset ServiceInstance at the specified address
func (group *ServiceGroup) serviceIsLive(addr *Address) {
	group.forEach(func(index int, instance *ServiceInstance) {
		if instance.ip == addr.IP && instance.port == addr.Port {
			instance.resetSurvivalTime()
		}
	})
}
