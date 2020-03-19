package trpc

import (
	"context"
	"log"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/proto"
)

// Register to Hub
func Register(serviceName string) error {
	config := configuration.GetConfig()

	providerInfo := proto.ProviderInfo{
		ServiceName: serviceName,
		Ip:          config.Self.Address,
		Port:        config.Self.Port,
	}

	conn, err := connentToHub()
	if err != nil {
		log.Fatalln(err)
	}

	client := proto.NewReceiverClient(conn)

	result, err := client.JoinServiceHub(context.Background(), &providerInfo)
	defer conn.Close()

	if result.Result {
		return nil
	}

	return err
}
