package client

import "net/rpc"

// ProviderInfo is RPC message format of Provider
type ProviderInfo struct {
	Name string
	IP   string
	Port string
}

// CustomerInfo is RPC message format of Customer
type CustomerInfo struct {
	ServiceName string
}

// Register to Hub
// TODO: Require TEST
func Register(serviceName string) error {
	config := GetConfig()

	client, err := rpc.DialHTTP("tcp", config.Hub.Address+":"+config.Hub.Port)

	if err != nil {
		return err
	}

	providerInfo := ProviderInfo{
		Name: serviceName,
		Port: config.Self.Port,
		IP:   config.Hub.Address,
	}

	var reply bool
	divCall := client.Go(JoinServiceHub, providerInfo, &reply, nil)
	replyCall := <-divCall.Done

	if replyCall.Error != nil {
		return replyCall.Error
	}

	return nil
}
