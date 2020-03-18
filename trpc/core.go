package trpc

import (
	"net/rpc"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/provider"
)

// CustomerInfo is RPC message format of Customer
type CustomerInfo struct {
	ServiceName string
}

// Call is RPC CALL of go-tds
func Call(serviceName string, args interface{}, reply interface{}) error {

	address, err := getServiceAddressFromHub(serviceName)

	if err != nil {
		return err
	}

	client, err := connect(address)
	if err != nil {
		return err
	}

	err = client.Call(configuration.TCP, args, reply)
	if err != nil {
		return err
	}
	return nil
}

func getServiceAddressFromHub(serviceName string) (*provider.Address, error) {
	config := configuration.GetConfig()

	info := CustomerInfo{ServiceName: serviceName}

	var address *provider.Address

	client, err := rpc.DialHTTP(configuration.TCP, config.Hub.Address+":"+config.Hub.Port)
	if err != nil {
		return address, err
	}

	replyDone := client.Go(configuration.GetService, info, address, nil)
	replyCall := <-replyDone.Done

	if replyCall.Error != nil {
		return address, replyCall.Error
	}

	return address, nil
}

// Create To RPC Client of Real service
func connect(addr *provider.Address) (*rpc.Client, error) {
	client, err := rpc.DialHTTP(configuration.TCP, addr.IP+":"+addr.Port)
	if err != nil {
		return client, err
	}
	return client, nil
}
