package trpc

import (
	"context"
	"log"

	"github.com/sdttttt/go-tds/configuration"
	"github.com/sdttttt/go-tds/proto"
	"github.com/sdttttt/go-tds/utils"
)

// TODO: Require TEST.
func keepConnectReport(info *proto.ProviderInfo) {
	conn, err := connectToHub()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	client := proto.NewReceiverClient(conn)
	timer := utils.NewTimer()
	config := configuration.GetConfig()

	timer.AddJob(config.Self.SurvivalTime, reportLive(client, info))
	timer.Run()
}

func reportLive(client proto.ReceiverClient, info *proto.ProviderInfo) func() {
	return func() {
		client.ReportActive(context.Background(), info)
	}
}
