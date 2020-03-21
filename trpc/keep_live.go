package trpc

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/sdttttt/go-tds/proto"
	"log"
)

func keepConnectReport(info *proto.ProviderInfo) {
	conn, err := connectToHub()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	client := proto.NewReceiverClient(conn)

	timer := cron.New()

	// TODO: CONSTANTS
	timer.AddFunc("45 * * * * * *", ReportLive(client, info))
	timer.Run()
}

func ReportLive(client proto.ReceiverClient, info *proto.ProviderInfo) func() {
	return func() {
		client.ReportActive(context.Background(), info)
	}
}
