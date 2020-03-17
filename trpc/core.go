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

	service, err := getServiceFromHub(serviceName)

	if err != nil {
		return err
	}

	client, err := connect(service)
	if err != nil {
		return err
	}

	err = client.Call(configuration.TCP, args, reply)
	if err != nil {
		return err
	}
	return nil
}

func getServiceFromHub(serviceName string) (*provider.Service, error) {
	config := configuration.GetConfig()

	info := CustomerInfo{ServiceName: serviceName}

	var service *provider.Service

	client, err := rpc.DialHTTP(configuration.TCP, config.Hub.Address+":"+config.Hub.Port)
	if err != nil {
		return service, err
	}

	replyDone := client.Go(configuration.GetService, info, service, nil)
	replyCall := <-replyDone.Done

	if replyCall.Error != nil {
		return service, replyCall.Error
	}

	return service, nil
}

func connect(service *provider.Service) (*rpc.Client, error) {
	client, err := rpc.DialHTTP(configuration.TCP, service.Address.IP+":"+service.Address.Port)
	if err != nil {
		return client, err
	}
	return client, nil
}
