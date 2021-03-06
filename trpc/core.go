package trpc

import (
	"github.com/sdttttt/go-tds/configuration"
	"google.golang.org/grpc"
)

func connect(ip string, port string) (*grpc.ClientConn, error) {
	return grpc.Dial(ip+":"+port, grpc.WithInsecure())
}

func connectToHub() (*grpc.ClientConn, error) {
	config := configuration.GetConfig()
	return connect(config.Hub.Address, config.Hub.Port)
}

// func connectToReceiver() (proto.ReceiverClient, error) {
// 	conn, err := connectToHub()
// 	return proto.NewReceiverClient(conn), err
// }

// func connectToEndPoint() (proto.EndPointClient, error) {
// 	conn, err := connectToHub()
// 	return proto.NewEndPointClient(conn), err
// }
