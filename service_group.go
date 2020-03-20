package main

// ServiceGroup is The Same Service, put in this group.
type ServiceGroup struct {
	servicesAddr []*Address
	useLen       uint8
	index        uint8

	arithmetic Balance
}

// NewServiceGroup is Get ServiceGroup Instance.
func NewServiceGroup(len uint8, bala Balance) *ServiceGroup {
	return &ServiceGroup{
		servicesAddr: make([]*Address, len),
		useLen:       0,
		index:        0,

		arithmetic: bala,
	}
}

// add is Join to ServiceGroup
func (group *ServiceGroup) add(addr *Address) {
	group.servicesAddr = append(group.servicesAddr, addr)
	group.useLen += 1
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
	return group.servicesAddr[index]
}
