package client

import (
	"net/rpc"

	"github.com/sdttttt/go-tds/configuration"
)

// ProviderInfo is RPC message format of Provider
type ProviderInfo struct {
	Name string
	IP   string
	Port string
}

// Register to Hub
// TODO: Require TEST
func Register(serviceName string) error {

	config := configuration.GetConfig()
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
	divCall := client.Go(configuration.JoinServiceHub, providerInfo, &reply, nil)
	replyCall := <-divCall.Done

	if replyCall.Error != nil {
		return replyCall.Error
	}

	return nil
}
