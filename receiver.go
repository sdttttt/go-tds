package main

import "github.com/sdttttt/go-tds/msg"

// Receiver is Rpc Receiver
type Receiver struct {
	Hub *ServiceHub
}

func (r *Receiver) JoinServiceHub(info msg.ProviderInfo, result *bool) error {
	r.Hub.Join(info.Name)
	*result = true
	return nil
}
