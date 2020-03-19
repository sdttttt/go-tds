package trpc

import (
	"context"
	"log"
	"net/rpc"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/proto"
)

// Call of Trpc
func Call(serviceName string, in interface{}, out interface{}) error {
	info, err := GetServiceAddr(serviceName)

	if err != nil {
		return err
	}

	client, err := buildConnection(&info)

	if err != nil {
		return err
	}

	err = client.Call(serviceName, in, out)
	defer client.Close()
	return err
}

// GetServiceAddr is Get Service Addr from Hub.
func GetServiceAddr(serviceName string) (proto.ProviderInfo, error) {
	conn, err := connentToHub()
	var info *proto.ProviderInfo

	if err != nil {
		log.Fatalln(err)
		return *info, err
	}

	client := proto.NewEndPointClient(conn)
	info, err = client.GetServiceInfo(context.Background(), &proto.ProviderInfo{ServiceName: serviceName})

	defer conn.Close()
	return *info, err
}

func buildConnection(info *proto.ProviderInfo) (*rpc.Client, error) {
	return rpc.DialHTTP(configuration.TCP, info.Ip+":"+info.Port)
}
